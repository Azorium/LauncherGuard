package main

import (
	"LauncherGuard/utils"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type FileWatcher struct{}

var fileUtils = utils.FileUtils{}

func (f *FileWatcher) run(gameDir string, cmd *exec.Cmd) {
	watcher, _ := fsnotify.NewWatcher()

	modsPath := filepath.Join(gameDir, "mods")
	librariesPath := filepath.Join(gameDir, "libraries")

	if fileUtils.Exist(modsPath) {
		filepath.Walk(modsPath, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				watcher.Add(path)
			}
			return nil
		})
	}

	if fileUtils.Exist(librariesPath) {
		filepath.Walk(librariesPath, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				watcher.Add(path)
			}
			return nil
		})
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Has(fsnotify.Write) {
					cmd.Process.Kill()
					os.Exit(0)
				}
				if event.Has(fsnotify.Create) {
					cmd.Process.Kill()
					os.Exit(0)
				}
				if event.Has(fsnotify.Remove) {
					cmd.Process.Kill()
					os.Exit(0)
				}
				if event.Has(fsnotify.Rename) {
					cmd.Process.Kill()
					os.Exit(0)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					fmt.Println(err.Error())
				}
				fmt.Println("Error:", err)
			}
		}

	}()
}
