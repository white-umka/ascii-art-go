package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	conv "ascii-art-go/converter"
)


func main() {
	width, height, img := conv.InputFile()
	result := conv.Converter(width, height, img) 

	step := 6
	for y := 0; y < height; y += step {
		for x := 0; x < width; x += step {
			fmt.Print(result[y][x])
		}
		fmt.Println()
	}
}