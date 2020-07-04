package mycaptcha

import (
	"bytes"
	"encoding/base64"
	"errors"


	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"math"
	"os"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

func init(){
	ReadFonts()
}

const defaultSaveFolder = "mycaptcha/image/"


// CaptchaImage    text is the captcha writed in the image
type CaptchaImage struct {
	rgba          *image.RGBA
	width, height int
	text          string
}

// NewCaptchaImage  initialization of the captchaimage struct
func NewCaptchaImage(width, height, n int, bgColor color.RGBA) *CaptchaImage {
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(m, m.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)
	return &CaptchaImage{
		rgba:   m,
		height: height,
		width:  width,
		text:   GetRandText(n),
	}
}

// GetText return CaptchaImage's text
func (c *CaptchaImage)GetText() string{
	return c.text
}

// SaveImage save CaptchaImage to saveFolder in png mode ,jpeg mode or gif mode and the filename is CaptchaImage.text
func (c *CaptchaImage) SaveImage(saveFolder string, imageFormat int)(pth string,err error ){
	if len(saveFolder)==0{
		saveFolder=defaultSaveFolder
	}
	pth=saveFolder+string(os.PathSeparator)+c.text
	switch imageFormat {
	case 1:
		pth+=".jpeg"
		file, err := os.Create(pth)
		if err != nil {
			return "",err
		}
		err=jpeg.Encode(file, c.rgba, nil)
		if err!=nil{
			return "",err
		}
	case 2:
		pth+=".gif"
		file, err := os.Create(pth)
		if err != nil {
			return "",err
		}
		err=gif.Encode(file, c.rgba, &gif.Options{NumColors: 256})
		if err!=nil{
			return "",err
		}
	default:
		pth+=".png"
		file, err := os.Create(pth)
		if err != nil {
			return "",err
		}
		err=png.Encode(file, c.rgba)
		if err!=nil{
			return "",err
		}
	}
	return pth,nil
}

// ImageToBaseb64 encode image to base64 or base64url mode 
func (c *CaptchaImage) ImageToBaseb64(format int) string {
	var buff bytes.Buffer
	var encodedString string
	png.Encode(&buff, c.rgba)
	if format ==0{
		encodedString = base64.StdEncoding.EncodeToString(buff.Bytes())
	}else{
		encodedString = base64.URLEncoding.EncodeToString(buff.Bytes())   

	}
	//htmlImage := "<img src=\"data:image/png;base64," + encodedString + "\" />"
	return encodedString
}

// DrawSinLine Draw a sine curve
func (c *CaptchaImage) DrawSinLine() {
	start := c.width / 20
	end := c.width - start
	lineColor := GetRandDeepColor()
	x1 := float64(r.Intn(start))
	x2 := float64(r.Intn(start) + end)
	w := c.height / 20

	yoffset := float64(c.height/4 + r.Intn(c.height/4))
	xoffset := 2 * r.Float64()
	for x1 < x2 {
		y := math.Sin(xoffset*math.Pi+x1*math.Pi/float64(c.width))*float64(c.height/4) + yoffset
		c.rgba.Set(int(x1), int(y), lineColor)
		for i := 1; i <= w; i++ {
			c.rgba.Set(int(x1), int(y)+i, lineColor)
		}
		x1++
	}
}

// DrawNoise draw noise
func (c *CaptchaImage) DrawNoise() {
	count := c.height * c.width / 50
	for i := 0; i < count; i++ {
		x := r.Intn(c.width)
		y := r.Intn(c.height)			
		c.rgba.Set(x, y,GetRandColor())

	}
}

// DrawCirlce maxsize : the max redius the one circle may have ;count: the number of circle
// The algorithm is parametric equations that x=r*sin(a) y=r*cos(a)
func (c *CaptchaImage) DrawCirlce(maxsize int,count int) {
	u:=math.Pi/360
	for i := 0; i < count; i++ {
		radius:=float64(1+r.Intn(maxsize))
		// center 
		cx := 5+r.Intn(c.width-5)
		cy := 5+r.Intn(c.height-5)
		for radius>0{
			color:=GetRandColor()
			for i:=0.0;i<720;i++{
				x:=cx+int(radius*math.Sin(i*u)+0.5)
				y:=cy+int(radius*math.Cos(i*u)+0.5)
				c.rgba.Set(x, y,color)
				
			}
			radius--
		}
	}
}

// DrawText  write the test on image
func (c *CaptchaImage) DrawText() error {
	//fmt.Println("text",c.text)
	ctx := freetype.NewContext()
	ctx.SetDPI(72)
	ctx.SetClip(c.rgba.Bounds())
	ctx.SetDst(c.rgba)
	ctx.SetHinting(font.HintingFull)
	fontWidth := c.width / len(c.text)
	for k, v := range c.text {
		fontSize := float64(c.height) * float64(10+r.Intn(4)) / 16
		ctx.SetSrc(image.NewUniform(GetRandDeepColor()))
		ctx.SetFontSize(fontSize)
		txfont, err :=GetRandFont()
		if err != nil {
			return errors.New("wrong in GetRandFont")
		}
		ctx.SetFont(txfont)
		x := fontWidth*k + fontWidth*(1+r.Intn(5))/20
		y := r.Intn(int(fontSize/3)) + c.height/5*3
		// x,y 表示字的左下角
		pt := freetype.Pt(x, y)
		_, err = ctx.DrawString(string(v), pt)
		if err != nil {
			return errors.New(" wrong in DrawString")
		}

	}
	return nil
}

// GetCaptchaBase64 This shortcut function contains several functions for generating captchaimage. 
//You can directly generate a captchaimage that encoded  base64(or base64url) mode and its corresponding captcha through this function
func GetCaptchaBase64(width, height, n int) (base64Str, text string) {
	CaptchaImage := NewCaptchaImage(width, height, n, GetRandLightColor())
	CaptchaImage.DrawSinLine()
	CaptchaImage.DrawNoise()
	CaptchaImage.DrawCirlce(height/5,width/height*3)
	CaptchaImage.DrawText()
	CaptchaImage.DrawCirlce(width/100,100)
	base64Str = CaptchaImage.ImageToBaseb64(0)
	text = CaptchaImage.text
	return
}


// GetSingleCaptcha This shortcut function contains several functions for generating captchaimage. 
//You can directly generate a captchaimage  and  save it through this function
// It's return  stored Image's path and  Captcha text
func GetSingleCaptcha(width, height, n int) (pth, text string) {
	CaptchaImage := NewCaptchaImage(width, height, n, GetRandLightColor())
	CaptchaImage.DrawNoise()
	CaptchaImage.DrawSinLine()
	CaptchaImage.DrawCirlce(height/3,width/height*3)
	CaptchaImage.DrawText()
	CaptchaImage.DrawCirlce(width/100,100)
	pth,err := CaptchaImage.SaveImage(".", 0)
	if err!=nil{
		return
	}
	text = CaptchaImage.text
	return pth,text

}
