package main

import (
	"os/exec"
	"fmt"
	"bytes"
)

const scale = 0.75
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

	// cmd = exec.Command("rm", "wallpaper_scaled_cropped.png")
	// if err := cmd.Run(); err != nil {
	// 	fmt.Printf("could not remove old scaled cropped wallpaper: %+q\n", err)
	// }

	// // center crop
	// cropArg := fmt.Sprintf("crop=%f:%f:%f:%f", baseWidth, baseHeight, ((zoom - 1.0) / 2.0) * baseWidth, ((zoom - 1.0) / 2.0) * baseHeight)
	// cmd = exec.Command("ffmpeg", "-i", "wallpaper_scaled.png", "-vf", cropArg, "wallpaper_scaled_cropped.png")
	// var stderrCC bytes.Buffer
	// cmd.Stderr = &stderrCC
	// if err := cmd.Run(); err != nil {
	// 	fmt.Println(fmt.Sprint(err) + ": " + stderrCC.String())
	// 	return
	// }

	// remove the old icon
	cmd = exec.Command("rm", "0_0.png")
	cmd.Run()

	// Crop the image
	cmd = exec.Command("ffmpeg", "-i", "wallpaper_scaled.png", "-vf", "crop=150:150:80:180", "0_0.png")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}

	fmt.Printf("finished!\n")
}

