package resolve_captcha

import (
	"image"
	"log"
)

type Pixel struct {
	R int
	G int
	B int
	A int
}

func CaptchaResolver(filename string) string {
	img := imageToThreshold(filename)
	saveImage("captcha_grey.png", img)

	pixels, err := getPixels(img)
	if err != nil {
		log.Fatal(err)
	}

	totalX := len(pixels[0])
	totalY := len(pixels)

	boxes := [][]Pixel{}
	for x := 0; x < totalX-8; x++ {
		for y := 0; y < totalY-12; y++ {
			box := []Pixel{}
			for startY := y; startY < y+12; startY++ {
				for startX := x; startX < x+8; startX++ {
					box = append(box, pixels[startY][startX])
				}
			}
			boxes = append(boxes, box)
		}
	}

	txtCaptcha := ""
	for _, box := range boxes {
		if checkNumberZero(box) {
			txtCaptcha += "0"
		} else if checkNumberOne(box) {
			txtCaptcha += "1"
		} else if checkNumberTwo(box) {
			txtCaptcha += "2"
		} else if checkNumberThree(box) {
			txtCaptcha += "3"
		} else if checkNumberFour(box) {
			txtCaptcha += "4"
		} else if checkNumberFive(box) {
			txtCaptcha += "5"
		} else if checkNumberSeven(box) {
			txtCaptcha += "7"
		} else if checkNumberSix(box) {
			txtCaptcha += "6"
		} else if checkNumberEight(box) {
			txtCaptcha += "8"
		} else if checkNumberNine(box) {
			txtCaptcha += "9"
		}
	}

	log.Println("captcha: ", txtCaptcha)

	return txtCaptcha
}

func rgbaToPixel(r, g, b, a uint32) Pixel {
	return Pixel{
		int(r / 257),
		int(g / 257),
		int(b / 257),
		int(a / 257),
	}
}

func getPixels(img image.Image) ([][]Pixel, error) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	pixels := [][]Pixel{}
	for y := 0; y < height; y++ {
		row := []Pixel{}
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}
