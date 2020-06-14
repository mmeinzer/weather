package main

import (
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype"
	imgcat "github.com/martinlindhe/imgcat/lib"
	"golang.org/x/image/font"
)

func main() {
	fontfile := "./avenir-next.ttc"

	fontBytes, err := ioutil.ReadFile(fontfile)
	if err != nil {
		log.Println(err)
		return
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	// Initialize the context.
	fg, bg := image.NewUniform(color.RGBA{177, 97, 133, 255}), image.Transparent
	rgba := image.NewRGBA(image.Rect(0, 0, 640, 100))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(220)
	c.SetFont(f)
	c.SetFontSize(24)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	c.SetHinting(font.HintingNone)

	// Draw the text.
	pt := freetype.Pt(10, 10+int(c.PointToFixed(24)>>6))
	s := "Test &"
	_, err = c.DrawString(s, pt)
	if err != nil {
		log.Println(err)
		return
	}
	pt.Y += c.PointToFixed(24 * 1.5)

	imgcat.CatImage(rgba, os.Stdout)
}
