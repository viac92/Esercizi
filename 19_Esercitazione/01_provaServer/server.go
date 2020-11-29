package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"net/http"

	"github.com/holizz/terrapin"
)

func quadratoHendler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/quadrato.png" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	i := image.NewRGBA(image.Rect(0, 0, 700, 700))
	t := terrapin.NewTerrapin(i, terrapin.Position{300.0, 400.0})
	for i := 0; i < 4; i++ {
		t.Forward(300)
		t.Left(math.Pi / 2)
	}
	png.Encode(w, i)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/quadrato.png", quadratoHendler)

	fmt.Print("Starting server port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
