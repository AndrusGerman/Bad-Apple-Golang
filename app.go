package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	for frame := 1; frame <= 5966; frame++ {
		clear()           // solo limpia la pantalla
		printFrame(frame) // que frame imprimir
		time.Sleep(29 * time.Millisecond)
	}
}

func printFrame(frame int) {
	var fileName = fmt.Sprintf("res/BA%d.txt", frame) // con el frame crea una variable del nombre del archivo
	var frameContent = getFrame(fileName)             // guarda el contenido del frame
	fmt.Fprint(os.Stdout, frameContent)               // imprime en pantalla el frame
}

func getFrame(frameDir string) string {
	fileRaw, _ := ioutil.ReadFile(frameDir) // lee el archivo
	return string(fileRaw)
}

func clear() {
	if runtime.GOOS == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	if runtime.GOOS == "windos" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
