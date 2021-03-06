package runner

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/howeyc/fsnotify"
)

func watchFolder(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fatal(err)
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if isWatchedFile(ev.Name) {
					startChannel <- ev.String()
				}
			case err := <-watcher.Error:
				watcherLog("Error: %s", err)
			}
		}
	}()

	err = watcher.Watch(path)

	if err != nil {
		fatal(err)
	}
}

func watch() {
	root := root()
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if len(path) > 1 && strings.HasPrefix(filepath.Base(path), ".") {
			return filepath.SkipDir
		}

		watchFolder(path)

		return err
	})
}
