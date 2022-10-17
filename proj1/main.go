package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)

		Area1 := r.FormValue("Area1")
		pincode1 := r.FormValue("Pincode1")
		fmt.Fprintf(w, "Area1       = %s\n", Area1)
		fmt.Fprintf(w, "Pincode1    = %s\n", pincode1)

		Area2 := r.FormValue("Area2")
		pincode2 := r.FormValue("Pincode2")
		fmt.Fprintf(w, "Area2       = %s\n", Area2)
		fmt.Fprintf(w, "Pincode1    = %s\n", pincode2)

		Area3 := r.FormValue("Area3")
		pincode3 := r.FormValue("Pincode3")
		fmt.Fprintf(w, "Area3      = %s\n", Area3)
		fmt.Fprintf(w, "Pincode3    = %s\n", pincode3)

		Area4 := r.FormValue("Area4")
		pincode4 := r.FormValue("Pincode4")
		fmt.Fprintf(w, "Area4       = %s\n", Area4)
		fmt.Fprintf(w, "Pincode4    = %s\n", pincode4)

		fmt.Fprintln(w, "...........................................")

		code := r.FormValue("code")

		if code == pincode1 {
			fmt.Fprintln(w, "Area name is =  ", Area1)

		} else if code == pincode2 {
			fmt.Fprintln(w, "Area name is = ", Area2)
		} else if code == pincode3 {
			fmt.Fprintln(w, "Area name is = ", Area3)
		} else if code == pincode4 {
			fmt.Fprintln(w, "Area name is =  ", Area4)
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST 8080 ...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
