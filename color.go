package mycaptcha

import (

	"image/color"
	"math/rand"
	"time"
)
var r=rand.New(rand.NewSource(time.Now().Unix()))

// We can get brightness of an RGB (or RGBA) value by calculate  r*0.299 + g*0.578 + b*0.114 
// In general if the result biggger than 192 ,it's a deep color 
 
// GetRandColor get a random  rgba color 
func GetRandColor()color.RGBA{
	c:=color.RGBA{}
	c.R=uint8(r.Intn(255))
	c.G=uint8(r.Intn(255))
	c.B=uint8(r.Intn(255))
	c.A=255
	//fmt.Println("GetRandColor",c)
	return c
}

// GetRandLightColor Get a random light color 
func GetRandLightColor()color.RGBA{
	colorR:=r.Intn(255)
	colorB:=r.Intn(255)
	tempG:=(185000-colorR*299-colorB*114)/578
	for tempG>254{
		if colorR>colorB{
			colorB=min(colorB+10,254)
		}else{
			colorR=min(colorR+10,254)
		}
		tempG=(185000-colorR*299-colorB*114)/578
	}

	colorG:=tempG+r.Intn(255-tempG)
	return color.RGBA{R: uint8(colorR),G: uint8(colorG),B: uint8(colorB),A: 255}
}

//GetRandDeepColor get a random deep color
func GetRandDeepColor()color.RGBA{
	colorR:=r.Intn(255)
	colorB:=r.Intn(255)
	tempG:=(210000-colorR*299-colorB*114)/578
	colorG:=r.Intn(tempG)
	return color.RGBA{R: uint8(colorR),G: uint8(colorG),B: uint8(colorB),A: 255}
}


func min(x,y int)int{
	if x<y{
		return x
	}
	return y
}