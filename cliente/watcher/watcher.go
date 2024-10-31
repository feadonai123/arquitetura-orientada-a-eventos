package watcher

import (
	"fmt"
	"strings"
	"time"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	utils "cliente/utils"
)

const WHO = "watcher"

type FileHandler func(content string)

func Run(
	dirInbox string,
	dirRead string,
	dirError string,
	fileHandlers map[string]FileHandler,
) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		utils.FailOnError(err, "Erro ao criar watcher")
	}
	defer watcher.Close()

	err = watcher.Add(dirInbox)
	if err != nil {
		utils.FailOnError(err, "Erro ao adicionar pasta ao watcher")
	}

	func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Op&fsnotify.Create == fsnotify.Create {
					utils.LogInfo(fmt.Sprintf("Novo arquivo detectado: %s", event.Name), WHO)

					time.Sleep(100 * time.Millisecond)

					content := utils.ReadFile(event.Name)

					isValid := false
					for key, handler := range fileHandlers {
						if strings.Contains(event.Name, key) {
							handler(string(content))
							isValid = true
							break
						}
					}

					if !isValid {
						utils.LogInfo(fmt.Sprintf("Arquivo %s não é válido", event.Name), WHO)
					}

					destErrorOrRead := dirError
					if isValid {
						destErrorOrRead = dirRead
					}

					destPath := filepath.Join(destErrorOrRead, filepath.Base(event.Name))
					utils.MoveFile(event.Name, destPath)
					utils.LogInfo(fmt.Sprintf("Movendo arquivo %s para a pasta %s", event.Name, destErrorOrRead), WHO)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				utils.FailOnError(err, "Erro ao monitorar pasta")
			}
		}
	}()
}