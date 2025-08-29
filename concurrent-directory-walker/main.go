package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func getFiles(path string, ch chan string) {
	defer close(ch)
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			ch <- path
		}

		return nil
	})

	if err != nil {
		log.Printf("Error during file walk: %v\n", err)
	}
}

func main() {

	pathsCh := make(chan string)

	var wg sync.WaitGroup

	var mu sync.Mutex
	fileHashes := make(map[string]string)

	var input string
	fmt.Println("Enter the path of the directory you wish to walk:")
	fmt.Scan(&input)

	go getFiles(input, pathsCh)

	workerCount := 5

	for range workerCount {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for path := range pathsCh {
				file, err := os.Open(path)
				if err != nil {
					log.Printf("Error opening file %q: %v\n", path, err)
					continue
				}

				hasher := sha256.New()
				if _, err := io.Copy(hasher, file); err != nil {
					log.Printf("Error hashing file %q: %v\n", path, err)
					file.Close()
					continue
				}
				file.Close()

				hash := fmt.Sprintf("%x", hasher.Sum(nil))

				mu.Lock()
				fileHashes[path] = hash
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	for i, v := range fileHashes {
		fmt.Println(i, ": ", v)
	}

}
