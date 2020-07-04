package mycaptcha

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)


// FontFamily a slice stores the all fonts path
var FontFamily []string

const (
	dirpath = "mycaptcha/fonts"
	//dirpath = "fonts"
	pthSep=string(os.PathSeparator)
)


// ReadFonts  Read all fonts in the folder dirpath
func ReadFonts()(err error){
	dir,err:=ioutil.ReadDir(dirpath)
	if err!=nil{
		return err
	}
	for _,fi :=range dir {
		if fi.IsDir(){
			continue
		}
		if strings.HasSuffix(strings.ToLower(fi.Name()),".ttf"){
			FontFamily=append(FontFamily,dirpath+pthSep+fi.Name())
		}
	}
	return nil
}
// GetRandFont  Randomly return a font
func GetRandFont()(*truetype.Font,error){
	fontFile:=FontFamily[r.Intn(len(FontFamily))]
	fontBytes,err := ioutil.ReadFile(fontFile)
	if err!=nil{
		return &truetype.Font{},err

	}
	f,err := freetype.ParseFont(fontBytes)
	if err!=nil{
		return &truetype.Font{},err

	}
	return f, nil
}
// GetFont Returns the font by fontName
func GetFont(fontName string)(*truetype.Font,error){
	fontBytes,err:=ioutil.ReadFile(dirpath+pthSep+fontName)
	if err!=nil{
		return &truetype.Font{},err

	}
	f,err := freetype.ParseFont(fontBytes)
	if err!=nil{
		return &truetype.Font{},err

	}
	return f, nil
}