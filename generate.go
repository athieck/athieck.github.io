package main

import (
	"os/exec"
	"fmt"
	"bytes"
)

const scale = 0.80
const baseWidth = 1242.0
const baseHeight = 2688.0

func main() {

	// remove the old scaled wallpaper
	cmd := exec.Command("rm", "wallpaper_scaled.png")
	cmd.Run()

	// scale the wallpaper
	cmd = exec.Command("ffmpeg", "-i", "wallpaper2.png", "-vf", fmt.Sprintf("scale=%f:%f", scale * baseWidth, scale * baseHeight), "wallpaper_scaled.png")
	if err := cmd.Run(); err != nil {
		fmt.Printf("could not scale wallpaper: %+q\n", err)
		return
	}

	for j := 0; j < 6; j++ {
		for i := 0; i < 4; i++ {
			fileName := fmt.Sprintf("%d_%d.png", j, i)

			// remove the old icon
			cmd = exec.Command("rm", fileName)
			cmd.Run()

			// Crop the image
			x := ((i+1) * 76.0) + (i * 150.0)
			y := 192.0 + (j * 150.0) + (j * 122.0)
			cropArg := fmt.Sprintf("crop=150:150:%d:%d", x, y)
			cmd = exec.Command("ffmpeg", "-i", "wallpaper_scaled.png", "-vf", cropArg, fileName)
			var stderr bytes.Buffer
			cmd.Stderr = &stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
				return
			}
		}
	}

	fmt.Printf("finished!\n")
}

