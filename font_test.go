package mycaptcha

import (
	"fmt"
	"testing"

)

func TestCaptchaFont_ReadFonts(t *testing.T) {
	
	ReadFonts()
	fmt.Println(FontFamily)
}

func TestCaptchaFont_GetRandFont(t *testing.T) {
	ReadFonts()
	fmt.Println(GetFont("actionj.ttf"))
}