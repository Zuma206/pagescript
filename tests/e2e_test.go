package psruntime

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/Zuma206/pagescript/psruntime"
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
	outputPath := getOutputPath(inputPath)
	expectedOutput, err := os.Open(outputPath)
	if err != nil {
		return fmt.Errorf("failed to open output file: %w", err)
	}
	output := new(bytes.Buffer)
	runtime := psruntime.NewPSRuntime()
	if err := runtime.Run(input, output); err != nil {
		return fmt.Errorf("failed to run input: %w", err)
	}
	if err := assertExpectedOutput(expectedOutput, output); err != nil {
		return err
	}
	return nil
}

func getOutputPath(inputPath string) string {
	inputPathWithoutSuffix, _ := strings.CutSuffix(inputPath, testSuffix)
	return inputPathWithoutSuffix + outputSuffix
}

func assertExpectedOutput(expectedOutput io.Reader, output io.Reader) error {
	bufferedExpectedOutput, err := io.ReadAll(expectedOutput)
	if err != nil {
		return fmt.Errorf("failed to real all of expected output: %w", err)
	}
	bufferedOutput, err := io.ReadAll(output)
	if err != nil {
		return fmt.Errorf("failed to real all of actual output: %w", err)
	}
	if !bytes.Equal(bufferedExpectedOutput, bufferedOutput) {
		return fmt.Errorf("actual output:\n%q\ndoesn't match expected output:\n%q", bufferedOutput, bufferedExpectedOutput)
	}
	return nil
}
