package converter

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
)


func Converter(width int, height int, img image.Image) [][]string {

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

	// создание матрицы для результата
	result := make([][]string, height)
	for i := range result {
		result[i] = make([]string, width)
	}

	// заполнение матрицы
	for y := range brightness {
		for x := range brightness[y] {
			result[y][x] = GetCharForBrightness(brightness[y][x])
		}
	}

	return result
}