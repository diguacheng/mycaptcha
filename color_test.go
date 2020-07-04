package mycaptcha

import (
	"image"
	"image/draw"
	"image/png"
	"os"
	"testing"
)

func TestGetRandColor(t *testing.T) {
	myimage := image.NewRGBA64(image.Rect(0, 0, 150, 50))
	c := GetRandColor()
	draw.Draw(myimage, myimage.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	file, err := os.Create("text.png")
	if err != nil {
		return
	}
	png.Encode(file, myimage)

}

func TestGetRandLightColor(t *testing.T) {
	myimage := image.NewRGBA64(image.Rect(0, 0, 150, 100))
	c := GetRandLightColor()
	c2:=GetRandDeepColor()
	draw.Draw(myimage, image.Rectangle{Min: image.Point{},Max: image.Point{X: 150, Y: 50}}, &image.Uniform{C: c}, image.ZP, draw.Src)
	draw.Draw(myimage, image.Rectangle{Min: image.Point{Y: 50},Max: image.Point{X: 150, Y: 100}}, &image.Uniform{C: c2}, image.ZP, draw.Src)
	file, err := os.Create("text.png")
	if err != nil {
		return
	}
	png.Encode(file, myimage)
	
}

func TestGetRandDeepColor(t *testing.T) {
	myimage := image.NewRGBA64(image.Rect(0, 0, 150, 50))
	c := GetRandDeepColor()
	draw.Draw(myimage, myimage.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	file, err := os.Create("text.png")
	if err != nil {
		return
	}
	png.Encode(file, myimage)
}