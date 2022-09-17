package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("NewWatcher failed", err)
		return
	}
	defer func(watcher *fsnotify.Watcher) {
		err := watcher.Close()
		if err != nil {
			log.Fatal("Close failed", err)
		}
	}(watcher)

	done := make(chan bool)
	go func() {
		defer close(done)

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Printf("%s %s\n", event.Name, event.Op)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println(err)
			}
		}
	}()
	err = watcher.Add("./")
	if err != nil {
		log.Fatal("Add failed", err)
	}
	<-done
}
