package main

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

// https://github.com/fsnotify/fsnotify
func main() {
	fileToWatch := os.Getenv("FILE_TO_WATCH")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	//resolvedFileToWatch, _ := filepath.EvalSymlinks(fileToWatch)
	err = watcher.Add(fileToWatch)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
