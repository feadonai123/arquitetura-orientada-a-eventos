package utils

import (
	"fmt"
	"os"
	"io/ioutil"
)

func MoveFile(src string, dst string) {
	err := os.Rename(src, dst)
	FailOnError(err, fmt.Sprintf("Erro ao mover arquivo %s para %s", src, dst))
}

func WriteFile(content string, dest string) {
	err := ioutil.WriteFile(dest, []byte(content), 0644)
	FailOnError(err, fmt.Sprintf("Erro ao escrever arquivo %s", dest))
}

func ReadFile(src string) string {
	content, err := ioutil.ReadFile(src)
	FailOnError(err, fmt.Sprintf("Erro ao ler arquivo %s", src))
	return string(content)
}