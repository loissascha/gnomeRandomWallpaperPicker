package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {
	startPath, err := os.UserHomeDir()
	if err != nil {
		panic("cant get user dir")
	}
	startPath += "/Pictures/randomWallpaper"

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for true {
			result := getAllFilesFromDir(startPath)
			randomIndex := rand.Intn(len(result))
			fmt.Println(result[randomIndex])
			time.Sleep(2 * time.Second)
		}
	}()

	fmt.Println("Waiting for goroutine to finish...")
	wg.Wait()
	fmt.Println("All goroutines finished!")

}

func getAllFilesFromDir(path string) []string {
	dir, err := os.ReadDir(path)
	if err != nil {
		return []string{}
	}

	result := []string{}
	for _, v := range dir {
		if v.IsDir() {
			result = append(result, getAllFilesFromDir(path+"/"+v.Name())...)
			continue
		}
		result = append(result, path+"/"+v.Name())
	}
	return result
}
