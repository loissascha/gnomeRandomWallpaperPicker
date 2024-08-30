package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func main() {
	var startPathFlag string
	var durationFlag int

	flag.StringVar(&startPathFlag, "path", "", "The path where random wallpapers should be picked from")
	flag.IntVar(&durationFlag, "duration", 300, "The duration in which the wallpaper should change. Default is 300 seconds.")

	flag.Parse()

	if startPathFlag == "" {
		fmt.Println("Please provide a -path flag!")
		return
	}

	fmt.Println(startPathFlag)
	fmt.Println(durationFlag)

	userDir, err := os.UserHomeDir()
	if err != nil {
		panic("cant get user dir")
	}

	for strings.Contains(startPathFlag, "~") {
		startPathFlag = strings.Replace(startPathFlag, "~", userDir, 1)
	}

	startPathFlag = strings.TrimSuffix(startPathFlag, "/")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for true {
			result := getAllFilesFromDir(startPathFlag)
			if len(result) == 0 {
				time.Sleep(time.Duration(durationFlag) * time.Second)
				continue
			}
			randomIndex := rand.Intn(len(result))
			fmt.Println("Next background:", result[randomIndex])
			cmd := exec.Command("bash", "-c", "gsettings set org.gnome.desktop.background picture-uri-dark \"file://"+result[randomIndex]+"\"")
			cmd.Run()
			cmd = exec.Command("bash", "-c", "gsettings set org.gnome.desktop.background picture-uri \"file://"+result[randomIndex]+"\"")
			cmd.Run()
			time.Sleep(time.Duration(durationFlag) * time.Second)
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
