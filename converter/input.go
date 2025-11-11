package converter

import (
	"fmt"
	"log"
	"image"
	"os"
	
)

func InputFile() (int, int, image.Image) {
	// получепние изображения
	var filePath string

	fmt.Println("enter file path")
	_, err := fmt.Scanln(&filePath)
	if err != nil {
		log.Fatal("incorrect file path:", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("error open file:", err)
	}
	defer file.Close()

	// декодируем изображение
	img, format, err := image.Decode(file)
	if err != nil {
		log.Fatal("error format image:", err)
	}
	fmt.Println("image format:", format)

	// определение размеров изображения
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	fmt.Println("x and y:", width, height)

	return width, height, img
}
