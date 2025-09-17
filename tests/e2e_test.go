package psruntime

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/Zuma206/pagescript/psruntime"
	"github.com/Zuma206/pagescript/stdlib"
)

const (
	dirPath      = "testdata"
	testSuffix   = ".html"
	outputSuffix = ".output"
)

func TestE2E(t *testing.T) {
	if err := testE2E(t); err != nil {
		t.Error(err.Error())
	}
}

func testE2E(t *testing.T) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("failed to read directory %q: %w", dirPath, err)
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), testSuffix) {
			continue
		}
		if err := runE2ETest(path.Join(dirPath, entry.Name())); err != nil {
			t.Errorf("e2e test %q failed: %s", entry.Name(), err.Error())
		}
	}
	return nil
}

func runE2ETest(inputPath string) error {
	input, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open test file: %w", err)
	}
	output, err := runTestInput(input)
	if err != nil {
		return fmt.Errorf("error running test: %w", err)
	}

	testName := strings.TrimSuffix(inputPath, testSuffix)
	expectedOutput, err := os.Open(testName + outputSuffix)
	if err != nil {
		return fmt.Errorf("failed to open expected output: %w", err)
	}

	return errors.Join(
		assertExpectedOutput("output", expectedOutput, output),
	)
}

func runTestInput(input io.Reader) (io.Reader, error) {
	var output bytes.Buffer

	runtime := psruntime.NewPSRuntime()
	stdlib.Open(runtime)
	go runtime.Eventloop().Start()

	if err := runtime.Run(input, &output); err != nil {
		return nil, fmt.Errorf("failed to run input: %w", err)
	}
	if err := runtime.Eventloop().Stop(); err != nil {
		return nil, fmt.Errorf("eventloop error: %w", err)
	}
	return &output, nil
}

func getOutputPath(inputPath string) string {
	inputPathWithoutSuffix, _ := strings.CutSuffix(inputPath, testSuffix)
	return inputPathWithoutSuffix + outputSuffix
}

func assertExpectedOutput(name string, expectedOutput io.Reader, output io.Reader) error {
	bufferedExpectedOutput, err := io.ReadAll(expectedOutput)
	if err != nil {
		return fmt.Errorf("failed to real all of expected %s: %w", name, err)
	}
	bufferedOutput, err := io.ReadAll(output)
	if err != nil {
		return fmt.Errorf("failed to real all of actual %s: %w", name, err)
	}
	if !bytes.Equal(bufferedExpectedOutput, bufferedOutput) {
		return fmt.Errorf("actual %s:\n%q\ndoesn't match expected %s:\n%q", name, bufferedOutput, name, bufferedExpectedOutput)
	}
	return nil
}
