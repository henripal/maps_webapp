package main

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"net/http"
	"strconv"
)

const apiKey string = "AIzaSyBMqdtriyHGTZHOsD2x-EJzqsK3N9PlPC4"

type subImager interface {
	SubImage(r image.Rectangle) image.Image
}

func cropSquare(img image.Image, size int) image.Image {
	return img.(subImager).SubImage(image.Rect(0, 0, size, size))
}

func getImage(location string, w int, h int, zoom int) (image.Image, error) {
	URL := makeURLFromParams(location, w, h, zoom)
	response, err := http.Get(URL)
	handleErr(err)

	return readImageFromResponse(response)

}

func readImageFromResponse(response *http.Response) (image.Image, error) {
	defer response.Body.Close()
	img, _, err := image.Decode(response.Body)
	return img, err
}

func makeURLFromParams(location string, w int, h int, zoom int) string {
	baseURL := "https://maps.googleapis.com/maps/api/staticmap?"
	sizeString := strconv.Itoa(w) + "x" + strconv.Itoa(h)
	return baseURL + "center=" + location +
		"&zoom=" + strconv.Itoa(zoom) +
		"&size=" + sizeString +
		"&maptype=satellite" +
		"&key=" + apiKey

}

func search(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad String", http.StatusBadRequest)
	}
	img, err := getImage(r.FormValue("location"), 224, 248, 13)
	img = cropSquare(img, 224)
	writeImage(w, &img)
}

func writeImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *img); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
