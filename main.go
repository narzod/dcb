package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
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
	Pattern  = os.Args[7]
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

func drawSpecialOutline(x, y, w, h int, pat string, col color.Color) {
	//bw := 5
	r := []rune(pat)
	for i, c := range r {
		//for i := 0; i <= bw-1; i++ {
		if c == '1' {
			drawRect(x+i, y+i, x+w-1-i, y+h-1-i, col)
			//fmt.Println(x+i, y+i, x+w-1-i, y+h-1-i, col)

		}
		//}
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

func getTilePhoto(url string) image.Image {
	fmt.Println("getting", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	image1, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// url := "https://i.picsum.photos/id/803/80/80.jpg?hmac=Z3s1SvDRg1QX9bC3auW_NL9rQakS6zZqlZoSIyJD0cU"
	// imageBytes := GetJpg(url)
	// fmt.Println(http.DetectContentType(imageBytes))
	tileimg, err := jpeg.Decode(bytes.NewReader(image1))
	if err != nil {
		panic("unable to decode jpeg")
	}
	return tileimg
}

func main() {

	var FgColor, BgColor, CurColor, AltColor color.Color

	// show args
	fmt.Println(os.Args[1:])

	if os.Args[2] == "random" {
		BgColor = x11colors.RandomSeeded().RGBA
	} else {
		BgColor = getRGBAByX11Name(os.Args[2])
	}

	if os.Args[1] == "random" {
		FgColor = x11colors.RandomSeeded().RGBA
	} else {
		FgColor = getRGBAByX11Name(os.Args[1])
	}

	MatrixWidth, MatrixHeight = parseAxB(os.Args[3])
	TileWidth, TileHeight = parseAxB(os.Args[4])

	// get jpeg from url and decode
	url2 := "http://www.fillmurray.com/" + fmt.Sprintf("%d", TileWidth) + "/" + fmt.Sprintf("%d", TileHeight)
	//url := "https://picsum.photos/" + fmt.Sprintf("%d", TileWidth) + "/" + fmt.Sprintf("%d", TileHeight)
	tileimg := getTilePhoto(url2)

	for col := 0; col <= MatrixHeight-1; col++ {
		for row := 0; row <= MatrixWidth-1; row++ {
			if isLightSquare(row, col) {
				CurColor = FgColor
				AltColor = BgColor
				//draw.Draw(img, tileimg.Bounds(), tileimg, image.Pt(0, 0), draw.Src)
				//drawTile(XOffset+(row*TileWidth), YOffset+(col*TileHeight), TileWidth, TileHeight, CurColor)
				topLtPoint := image.Pt(XOffset+(row*TileWidth), YOffset+(col*TileHeight))
				botRtPoint := image.Pt(XOffset+(row*TileWidth)+TileWidth, YOffset+(col*TileHeight)+TileHeight)
				tileRect := image.Rectangle{topLtPoint, botRtPoint}

				draw.Draw(img, tileRect, tileimg, image.Pt(0, 0), draw.Src)

			} else {
				CurColor = BgColor
				AltColor = FgColor
				drawTile(XOffset+(row*TileWidth), YOffset+(col*TileHeight), TileWidth, TileHeight, CurColor)
			}
			//drawTileOutline(XOffset+(row*TileWidth), YOffset+(col*TileHeight), TileWidth, TileHeight, BorderWidth, AltColor)
			drawSpecialOutline(XOffset+(row*TileWidth), YOffset+(col*TileHeight), TileWidth, TileHeight, Pattern, AltColor)
			// DrawTileOutline(XOffset+(row*TileWidth)+2, YOffset+(col*TileHeight)+2, TileWidth-4, TileHeight-4, AltColor)
			// DrawTileOutline(XOffset+(row*TileWidth)+4, YOffset+(col*TileHeight)+4, TileWidth-8, TileHeight-8, AltColor)
			//fmt.Println(BorderMultipler)
			//drawTileOutline(XOffset+(row*TileWidth)+10, YOffset+(col*TileHeight)+10, TileWidth-20, TileHeight-20, BorderWidth, AltColor)

		}
	}

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
