package main

import (
	"os"
	"image"
	"fmt"
	"log"
	_ "image/jpeg"
	_ "image/png"
)


func main() {
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

	bounds := img.Bounds()
	width := bounds.Max.X 
	height := bounds.Max.Y
	fmt.Println("x and y:", width, height)

	brightness := make([][]float64, height) // создай список с height строк
	for i := range brightness { // для каждой строки в этом списке
		brightness[i] = make([]float64, width) // создай внутри нее список из width чисел 
	}
	
	// проходимя по всем пикселям изображения 
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixelColor := img.At(x, y) // получаем цвет пикселя в координатах (x; y)

			r, g, b, _ := pixelColor.RGBA()

			// вычисление яркости пикселя, сохранение в матрицу
			brightness[y][x] = (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 65535 * 100
		}
	}
	result := make([][]string, height)
	for i := range result {
		result[i] = make([]string, width)
	}

	for y := range brightness {
		for x := range brightness[y] {
			result[y][x] = getCharForBrightness(brightness[y][x])
		}
	}

	step := 4
	for y := 0; y < height; y += step {
		for  x := 0; x < width; x += step {
			fmt.Print(result[y][x])
		}
		fmt.Println()
	}

}


func getCharForBrightness(brightness float64) string { // 1 - черный, 100 - белый
    switch {
    case brightness > 90:
        return "@"  
    case brightness > 80:
        return "0"  
    case brightness > 65:
        return "O"
    case brightness > 45:
        return "o"
    case brightness > 30:
        return "*"
    case brightness > 15:
        return "."
    default:
        return " "  
    }
}