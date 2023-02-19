package main

import (
	"fmt"
	"io"

	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

var framesChan = make(chan string, 1000)
var wg = new(sync.WaitGroup)

func main() {
	wg.Add(1)
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Load Frames
	go LoadFrames()

	go PrintFrames()

	wg.Wait()
}

func PrintFrames() {
	for frameContent := range framesChan {
		clear()
		printFrame(frameContent)
		time.Sleep(31 * time.Millisecond)
		wg.Done()
	}
}

func LoadFrames() {
	for frame := 1; frame <= 6569; frame++ {
		wg.Add(1)
		var fileName = fmt.Sprintf("res/BA%d.txt", frame) // con el frame crea una variable del nombre del archivo
		var frameContent = getFrameContent(fileName)
		framesChan <- frameContent
	}
	wg.Done()
}

func printFrame(frameContent string) {
	fmt.Fprint(os.Stdout, frameContent) // imprime en pantalla el frame
}

func getFrameContent(frameDir string) string {
	// Open File
	file, err := os.Open(frameDir)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read io.Reader
	fileRaw, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return string(fileRaw)
}

func clear() {
	if runtime.GOOS == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	if runtime.GOOS == "windos" {
		fmt.Print("\033[H\033[2J")
	}
}
