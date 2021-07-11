package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"net/http/httputil"
)

func GetJpg(url string) []byte {

	//url := "http://i.imgur.com/m1UIjW1.jpg"
	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	b, err := httputil.DumpResponse(response, true)
	if err != nil {
		log.Fatalln(err)
	}

	return b
}

// ToPng converts an image to png
func ToPng(imageBytes []byte) ([]byte, error) {
	contentType := http.DetectContentType(imageBytes)

	switch contentType {
	case "image/png":
	case "image/jpeg":
		img, err := jpeg.Decode(bytes.NewReader(imageBytes))
		if err != nil {
			return nil, err // errors.Wrap(err, "unable to decode jpeg")
		}

		buf := new(bytes.Buffer)
		if err := png.Encode(buf, img); err != nil {
			return nil, err // errors.Wrap(err, "unable to encode png")
		}

		return buf.Bytes(), nil
	}

	return nil, fmt.Errorf("unable to convert %#v to png", contentType)
}
