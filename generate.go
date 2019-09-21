package main

import (
	"os/exec"
	"fmt"
	"bytes"
)

func main() {
	cmd := exec.Command("ffmpeg", "-i", "wallpaper.png", "-vf", "crop=150:150:40:80", "0_0.png")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Printf("finished!\n")
}