package imageprocessing

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func ReadImage(path string) image.Image {
	inputFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	//Decode image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Println(path)
		panic(err)
	}
	return img
}

func WriteImage(path string, img image.Image) {
	outFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	//Encode the image to the new file
	err = jpeg.Encode(outFile, img, nil)
	if err != nil {
		panic(err)
	}
}

func GrayScale(img image.Image) image.Image {
	//Create a new grayscale image
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	//Convert each pixel to gray
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originPixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(originPixel)
			grayImg.Set(x, y, grayPixel)
		}
	}
	return grayImg
}

func Resize(img image.Image) image.Image {
	newWidth := uint(500)
	newHeight := uint(500)

	resizeImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos2)
	return resizeImg
}
