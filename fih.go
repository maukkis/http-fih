package main

import (
    "bytes"
    "errors"
    "fmt"
    "image"
    "image/jpeg"
    "log"
    "math/rand/v2"
    "net/http"
    "os"
    "strconv"
)

func main() {
    port := ":8080"

    http.HandleFunc("/fih", fihHandler)
    log.Println("Listening on:", port)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
        return
    }

    fihlist,_ := os.ReadDir("./fihhes")
    fishCount := (len(fihlist))
    fmt.Println(fishCount)
}

func getFih(fihDir string) (image.Image, error) {
    fihlist, _ := os.ReadDir(fihDir)
    fihCount := (len(fihlist))
    if fihCount <= 0 {
        return nil, errors.New("no fih")
    }
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
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, err.Error())
        return
    }
    writeImage(w, &img)
}


func writeImage(w http.ResponseWriter, img *image.Image) {
    buffer := new(bytes.Buffer)
    if err := jpeg.Encode(buffer, *img, nil); err != nil {
        log.Println("unable to encode image.")
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "something went wrong while encoding the image *tail wags*");
        return
    }

    w.Header().Set("Content-Type", "image/jpeg")
    w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
    if _, err := w.Write(buffer.Bytes()); err != nil {
        log.Println("unable to write image.")
    }
}
