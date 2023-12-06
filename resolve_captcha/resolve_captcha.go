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

func rgbaToPixel(r, g, b, a uint32) Pixel {
	return Pixel{
		int(r / 257),
		int(g / 257),
		int(b / 257),
		int(a / 257),
	}
}

func GetPixels(img image.Image) ([][]Pixel, error) {
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

func CaptchaResolver() string {
	img := LoadImageThreshold("captcha.png")

	pixels, err := GetPixels(img)
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
		} else if checkNumberTwo(box) {
			txtCaptcha += "2"
		} else if checkNumberFour(box) {
			txtCaptcha += "4"
		}
	}

	log.Println("captcha: ", txtCaptcha)

	return txtCaptcha
}

func checkNumberZero(pixel []Pixel) bool {
	var (
		numberTwo []bool
		px        Pixel
	)
	px = pixel[3]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[4]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[5]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[9]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[10]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[11]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[12]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[13]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[14]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[15]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[17]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[18]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[22]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[23]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[24]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[25]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[31]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[32]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[33]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[39]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[40]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[41]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[47]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[48]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[49]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[55]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[56]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[57]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[63]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[64]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[65]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[71]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[73]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[74]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[78]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[79]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[81]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[82]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[83]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[84]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[85]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[86]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[87]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[91]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[92]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[93]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)

	notNumberZero := false
	for _, two := range numberTwo {
		if !two {
			notNumberZero = true
		}
	}
	return !notNumberZero
}

func checkNumberTwo(pixel []Pixel) bool {
	var (
		numberTwo []bool
		px        Pixel
	)
	px = pixel[1]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[2]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[3]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[4]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[5]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[8]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[9]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[10]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[11]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[12]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[13]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[14]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[16]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[22]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[23]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[30]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[31]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[38]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[39]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[45]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[46]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[51]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[52]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[53]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[58]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[59]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[65]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[66]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[72]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[73]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[80]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[81]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[82]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[83]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[84]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[85]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[86]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[87]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[88]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[89]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[90]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[91]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[92]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[93]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[94]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[95]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)

	notNumberTwo := false
	for _, two := range numberTwo {
		if !two {
			notNumberTwo = true
		}
	}
	return !notNumberTwo
}

func checkNumberFour(pixel []Pixel) bool {
	var (
		numberTwo []bool
		px        Pixel
	)
	px = pixel[6]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[7]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[13]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[14]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[15]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[20]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[21]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[22]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[23]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[27]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[28]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[30]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[31]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[34]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[35]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[38]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[39]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[41]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[42]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[46]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[47]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[48]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[49]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[54]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[55]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[56]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[57]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[58]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[59]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[60]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[61]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[62]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[63]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[64]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[65]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[66]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[67]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[68]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[69]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[70]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[71]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[78]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[79]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[86]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[87]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[94]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)
	px = pixel[95]
	numberTwo = append(numberTwo, px.R == 0, px.G == 0, px.B == 0)

	notNumberFour := false
	for _, two := range numberTwo {
		if !two {
			notNumberFour = true
		}
	}
	return !notNumberFour
}
