package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
	"os"
	"fmt"
	"math/rand/v2"
	"io/ioutil"
)

func main() {
	port := ":8080"

	http.HandleFunc("/fih", fihHandler)
	log.Println("Listening on:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	fihlist,_ := ioutil.ReadDir("./fihhes")
	fishCount := (len(fihlist))
	fmt.Println(fishCount)

}

func getFih(fihDir string) (image.Image, error) {
	fihlist,_ := ioutil.ReadDir(fihDir)
	fihCount := (len(fihlist))
	fmt.Println(fihCount)
	fih := fmt.Sprintf("%s/fih%d.jpg",fihDir, rand.IntN(fihCount) )
	fmt.Println(fih)
	f, err := os.Open(fih)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

func fihHandler(w http.ResponseWriter, r *http.Request) {
	img, err := getFih("./fihhes")
	if err != nil{
		fmt.Fprintf(w, "fuck you\n")
	}
	writeImage(w, &img)
}


func writeImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
