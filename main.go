package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	img1, err := imageFrom("en61.jpg")
	if err != nil {
		return fmt.Errorf("image1 read error = %w", err)
	}
	img2, err := imageFrom("ja61.png")
	if err != nil {
		return fmt.Errorf("image2 read error = %w", err)
	}
	img := joinImage(img1, img2)
	out, err := os.Create("out2.png")
	if err != nil {
		return err
	}
	defer out.Close()
	return png.Encode(out, img)
}

func imageFrom(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func maxWidth(image1x, image2x int) int {
	if image1x > image2x {
		return image1x
	}
	return image2x
}

func joinImage(img1, img2 image.Image) image.Image {
	bounds1 := img1.Bounds()
	bounds2 := img2.Bounds()
	// If height > width, draw image on left side.
	width := maxWidth(bounds1.Dx(), bounds2.Dx())
	bounds := image.Rect(0, 0, width, bounds1.Dy()+bounds2.Dy()+2)
	img := image.NewRGBA(bounds)

	for h := 0; h < width; h++ {
		img.Set(h, bounds1.Dy()+1, color.RGBA{255, 255, 255, 255})
	}

	draw.Draw(img, bounds1, img1, image.ZP, draw.Src)
	draw.Draw(img, bounds2.Add(image.Pt(0, bounds1.Dy()+2)), img2, image.ZP, draw.Src)

	return img
}
