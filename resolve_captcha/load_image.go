package resolve_captcha

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/disintegration/gift"
)

func imageToThreshold(filename string) image.Image {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("os.Open failed: %v", err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("image.Decode failed: %v", err)
	}

	filter := gift.Threshold(50)
	g := gift.New(filter)
	dst := image.NewNRGBA(g.Bounds(img.Bounds()))
	g.Draw(dst, img)

	return dst
}

func saveImage(filename string, img image.Image) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("os.Create failed: %v", err)
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		log.Fatalf("png.Encode failed: %v", err)
	}
}
