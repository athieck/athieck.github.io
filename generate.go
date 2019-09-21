package main

import (
	"os/exec"
	"fmt"
	"bytes"
)

func main() {
	cmd := exec.Command("rm", "0_0.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("could not remove\n")
		return
	}
	cmd = exec.Command("ffmpeg", "-i", "wallpaper2.png", "-vf", "crop=150:150:160:270", "0_0.png")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Printf("finished!\n")
}