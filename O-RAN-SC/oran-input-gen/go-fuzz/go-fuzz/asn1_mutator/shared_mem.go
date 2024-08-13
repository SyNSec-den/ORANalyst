package asn1_mutator

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	mmap "github.com/edsrzf/mmap-go"
)

func mutateMessageFork(data []byte) ([]byte, error) {

	// Create a temp file for shared memory.
	sharedFile, err := os.CreateTemp("", "mutator-shared")
	if err != nil {
		return make([]byte, 0), err
	}
	defer os.Remove(sharedFile.Name())
	defer sharedFile.Close()

	// Resize the file to fit the data.
	if err := sharedFile.Truncate(2048 * 2048); err != nil {
		return make([]byte, 0), err
	}

	// Memory-map the file.
	mmap, err := mmap.Map(sharedFile, mmap.RDWR, 0)
	if err != nil {
		return make([]byte, 0), err
	}
	defer mmap.Unmap()

	// Write to shared memory.
	copy(mmap, data)

	// Create extra pipes for communication.
	pipeInReader, pipeInWriter, err := os.Pipe()
	if err != nil {
		return make([]byte, 0), err
	}
	pipeOutReader, pipeOutWriter, err := os.Pipe()
	if err != nil {
		return make([]byte, 0), err
	}

	// Invoke the subprocess, passing the file descriptors for the pipes.
	cmd := exec.Command("mutator/kpm/mutator", sharedFile.Name())
	cmd.ExtraFiles = []*os.File{sharedFile, pipeInReader, pipeOutWriter}

	// Create pipes for the subprocess's standard output and standard error.
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return make([]byte, 0), err
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return make([]byte, 0), err
	}

	// Start the subprocess.
	if err := cmd.Start(); err != nil {
		return make([]byte, 0), err
	}

	// Create goroutines to read and print the subprocess's stdout and stderr.
	go func() {
		io.Copy(os.Stdout, stdoutPipe)
	}()
	go func() {
		io.Copy(os.Stderr, stderrPipe)
	}()

	// We don't need the subprocess's end of the pipes in the main process.
	pipeInReader.Close()
	pipeOutWriter.Close()

	// Send the length of the data to the subprocess.
	fmt.Fprintln(pipeInWriter, len(data))
	pipeInWriter.Close() // Close the writer to signal we're done sending.

	// Read the response length from the subprocess.
	var responseLength int
	fmt.Fscanf(pipeOutReader, "%d", &responseLength)

	// Read the response data directly from the shared memory area.
	response := make([]byte, responseLength)
	copy(response, mmap[:responseLength])

	// Wait for the subprocess to finish.
	if err := cmd.Wait(); err != nil {
		return make([]byte, 0), err
	}

	if len(response) == 0 {
		return make([]byte, 0), fmt.Errorf("response is empty")
	}

	// fmt.Printf("Received response: %v\n", response)
	return response, nil
}
