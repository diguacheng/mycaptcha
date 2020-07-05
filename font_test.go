package mycaptcha

import (
	"fmt"
	"testing"

)

func TestCaptchaFont_ReadFonts(t *testing.T) {
	
	LoadFonts("foonts")
	fmt.Println(FontFamily)
}

func TestCaptchaFont_GetRandFont(t *testing.T) {
	LoadFonts("foonts")
	fmt.Println(GetFont("fonts","actionj.ttf"))
}