package mycaptcha

import (
	"fmt"
	"image/color"

	"reflect"
	"testing"
)

func TestCaptchaImage_DrawSinLine(t *testing.T) {
	captchaImage := NewCaptchaImage(150, 50,4, GetRandLightColor())
	captchaImage.DrawSinLine()
	pth,err := captchaImage.SaveImage("image", 0)
	fmt.Println(pth)
	if err != nil {
		return
	}

}

func TestCaptchaImage_DrawNoise(t *testing.T) {
	captchaImage := NewCaptchaImage(150, 50, 4,GetRandLightColor())
	captchaImage.DrawSinLine()
	captchaImage.DrawNoise()
	_,err:= captchaImage.SaveImage("image", 0)
	if err != nil {
		return
	}
}

func TestCaptchaImage_DrawText(t *testing.T) {
	captchaImage := NewCaptchaImage(150, 50, 4,GetRandLightColor())
	captchaImage.DrawSinLine()
	captchaImage.DrawNoise()
	err := captchaImage.DrawText()
	if err != nil {
		return
	}

	_,err = captchaImage.SaveImage("image", 0)
	if err != nil {
		return
	}
}

func TestNewCaptchaImage(t *testing.T) {
	type args struct {
		width   int
		height  int
		bgColor color.RGBA
	}
	tests := []struct {
		name string
		args args
		want *CaptchaImage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCaptchaImage(tt.args.width, tt.args.height,4, tt.args.bgColor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCaptchaImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCaptchaBase64(t *testing.T) {
	for i := 0; i < 10; i++ {
		_, s2 := GetCaptchaBase64(300, 100, 5)
		//fmt.Println(s1)
		fmt.Println(s2)
	}

}

func TestGetSingleCaptcha(t *testing.T) {
	
	for i := 0; i <1; i++ {
		//pth,text:=GetSingleCaptcha(600,200,5)
		_,_=GetSingleCaptcha(20000,1000,40)
		//fmt.Println(s1)
		//fmt.Println(pth,text)
	}

}
