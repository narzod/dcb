package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"

	"github.com/vgarvardt/x11colors-go"
)

var (
	//MatrixWidth  = 8
	//MatrixHeight = 4
	//TileWidth    = 80
	//TileHeight   = 80
	XOffset  = 0
	YOffset  = 0
	Color1   = "Cornsilk"
	Color2   = "Tomato"
	fileName = os.Args[6]
)

var TileWidth, TileHeight = parseAxB(os.Args[3])
var MatrixWidth, MatrixHeight = parseAxB(os.Args[4])
var BorderWidth, BorderMultipler = parseAxB(os.Args[5])

var img = image.NewRGBA(image.Rect(0, 0, TileWidth*MatrixWidth, TileHeight*MatrixHeight))

func parseAxB(s string) (int, int) {
	slice1 := strings.Split(s, "x")
	a, _ := strconv.Atoi(slice1[0])
	b, _ := strconv.Atoi(slice1[1])
	// really gotta do error checking!!
	return a, b
}

// hLine draws a horizontal line
func hLine(x1, y, x2 int, col color.Color) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// vLine draws a veritcal line
func vLine(x, y1, y2 int, col color.Color) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// drawRect draws a rectangle utilizing hLine() and vLine()
func drawRect(x1, y1, x2, y2 int, col color.Color) {
	hLine(x1, y1, x2, col)
	hLine(x1, y2, x2, col)
	vLine(x1, y1, y2, col)
	vLine(x2, y1, y2, col)
}

// drawSolidRect draws solid rectangle utilizing hLine()
func drawSolidRect(x1, y1, x2, y2 int, col color.Color) {
	for ; y1 <= y2; y1++ {
		hLine(x1, y1, x2, col)
	}
}

func drawTile(x, y, w, h int, col color.Color) {
	y1 := y
	for ; y <= y1+h-1; y++ {
		hLine(x, y, x+w-1, col)
	}
}

func drawTileOutline(x, y, w, h, bw int, col color.Color) {
	for i := 0; i <= bw-1; i++ {
		drawRect(x+i, y+i, x+w-1-i, y+h-1-i, col)
	}
}

// also see alt algorithm https://en.wikipedia.org/wiki/Checkerboard#Mathematical_description
func isLightSquare(x, y int) bool {
	if (x%2 == 0) && (y%2 == 0) {
		return true
	}
	if (x%2 != 0) && (y%2 != 0) {
		return true
	}
	return false
}

func getRGBAByX11Name(name string) color.RGBA {
	name = strings.Title(strings.ToLower(name))
	xc, _ := x11colors.GetByName(name)
	// fmt.Println(name, test)
	return xc.RGBA
}

func main() {

	var FgColor, BgColor, CurColor, AltColor color.Color

	fmt.Println(os.Args[1:])

	FgColor = color.RGBA{255, 0, 0, 255}   // Red
	BgColor = color.RGBA{255, 0, 255, 255} // Purple

	FgColor = getRGBAByX11Name(os.Args[1])
	BgColor = getRGBAByX11Name(os.Args[2])

	MatrixWidth, MatrixHeight = parseAxB(os.Args[3])
	fmt.Println(MatrixWidth, MatrixHeight)

	TileWidth, TileHeight = parseAxB(os.Args[4])

	for col := 0; col <= MatrixHeight-1; col++ {
		for row := 0; row <= MatrixWidth-1; row++ {
			if isLightSquare(row, col) {
				CurColor = FgColor
				AltColor = BgColor
			} else {
				CurColor = BgColor
				AltColor = FgColor
			}
			drawTile(XOffset+(row*TileWidth), YOffset+(col*TileHeight), TileWidth, TileHeight, CurColor)
			drawTileOutline(XOffset+(row*TileWidth), YOffset+(col*TileHeight), TileWidth, TileHeight, 5, AltColor)
			// DrawTileOutline(XOffset+(row*TileWidth)+2, YOffset+(col*TileHeight)+2, TileWidth-4, TileHeight-4, AltColor)
			// DrawTileOutline(XOffset+(row*TileWidth)+4, YOffset+(col*TileHeight)+4, TileWidth-8, TileHeight-8, AltColor)
			fmt.Println(BorderMultipler)
			drawTileOutline(XOffset+(row*TileWidth)+10, YOffset+(col*TileHeight)+10, TileWidth-20, TileHeight-20, BorderWidth, AltColor)

		}
	}

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
