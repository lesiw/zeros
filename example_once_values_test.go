package zeros_test

import (
	"fmt"
	"io"
	"os"

	"lesiw.io/zeros"
)

type lazyFile struct {
	once zeros.OnceValues[*os.File, error]
}

func (f *lazyFile) init() (*os.File, error) {
	fmt.Println("Creating temp file")
	return os.CreateTemp("", "example")
}

func (f *lazyFile) Write(p []byte) (int, error) {
	file, err := f.once.Do(f.init)
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %w", err)
	}
	return file.Write(p)
}

func (f *lazyFile) Close() error {
	file, err := f.once.Do(f.init)
	if err != nil { 
		return fmt.Errorf("failed to close file: %w", err)
	}
	defer os.Remove(file.Name())
	return file.Close()
}

func (f *lazyFile) Stat() (os.FileInfo, error) {
	file, err := f.once.Do(f.init)
	if err != nil {
		return nil, fmt.Errorf("failed to stat file: %w", err)
	}
	return file.Stat()
}

func ExampleOnceValues() {
	var f lazyFile
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		fmt.Printf("Stat failed: %v\n", err)
		return
	}
	fmt.Printf("Initial size: %d bytes\n", info.Size())

	if _, err := io.WriteString(&f, "Hello, world!"); err != nil {
		fmt.Printf("Copy failed: %v\n", err)
		return
	}

	info, err = f.Stat()
	if err != nil {
		fmt.Printf("Stat failed: %v\n", err)
		return
	}
	fmt.Printf("Final size: %d bytes\n", info.Size())

	// Output:
	// Creating temp file
	// Initial size: 0 bytes
	// Final size: 13 bytes
}
