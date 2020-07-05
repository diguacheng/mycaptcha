package main

import (
	//"fmt"
	"net/http"
	
	"github.com/diguacheng/mycaptcha"
	"github.com/gin-gonic/gin"


)

const fontfolder="fonts"

func init(){
	// 导入字体
	mycaptcha.LoadFonts(fontfolder)
}

func CustomCaptchaImage()(string, string){
	//w,h:=600,200
	CaptchaImage :=mycaptcha.NewCaptchaImage(150,50, 4, mycaptcha.GetRandLightColor())
	CaptchaImage.DrawNoise()
	CaptchaImage.DrawSinLine()
	
	CaptchaImage.DrawCircle(10,5)
	CaptchaImage.DrawText()
	
	CaptchaImage.DrawCircle(3,50)
	saveFolder:="."
	pth,err := CaptchaImage.SaveImage(saveFolder, 0)
	if err != nil {
		return  "",""
	}
	text := CaptchaImage.GetText()
	return pth,text
	


}

func main(){
	// base64 encoded image used in website
	r:=gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("index",func (c *gin.Context){
		base64image,_:=mycaptcha.GetCaptchaBase64(300,100,4)
		c.HTML( http.StatusOK,"index.html",gin.H{"base64image":base64image} )

	})
	r.Run(":8086")

	
	// Custom CaptchaImage by yourself
	//pth,text:=CustomCaptchaImage() 

	
	// one function get a captcha image and captcha text
	//pth,text:=mycaptcha.GetSingleCaptcha(300,100,4)

	// one function get a captcha image encoded by base64 and captcha test
	//base64str,test:=mycaptcha.GetCaptchaBase64(150,50,5)
	
	
}