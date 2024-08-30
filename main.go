package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
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
			fmt.Println("Next background:", result[randomIndex])
			cmd := exec.Command("bash", "-c", "gsettings set org.gnome.desktop.background picture-uri-dark \"file://"+result[randomIndex]+"\"")
			cmd.Run()
			cmd = exec.Command("bash", "-c", "gsettings set org.gnome.desktop.background picture-uri \"file://"+result[randomIndex]+"\"")
			cmd.Run()
			time.Sleep(300 * time.Second)
		}
	}()

	wg.Wait()
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
