package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// allowing clients to pass inputs inside request body

// creating an handler/endpoint that will serve videos data over a get request
func HandleGetVideos(w http.ResponseWriter, r *http.Request) {
	videos := getVideos()
	// coverting slice of videos slices into slice of bytes
	videoBytes, err := json.Marshal(videos)
	// handling error
	if err != nil {
		panic(err)
	}
	// returning videobytes into response writer
	w.Write(videoBytes)
}

func HandleUpdateVideos(w http.ResponseWriter, r *http.Request) {
	// validating if the method is POST or not
	if r.Method == "POST" {
		// getting and reading the request body into a slice of bytes
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		// converting slice of bytes into a slice of videos
		var videos []video
		err = json.Unmarshal(body, &videos)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprint(w, "Bad Request!")
		}
		// saving back the videos
		saveVideos(videos)
	} else {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method is not supported :-(")
	}
}

func main() {
	// endpoint to serve video data
	http.HandleFunc("/", HandleGetVideos)
	// endpoint to update users videos
	http.HandleFunc("/update", HandleUpdateVideos)
	http.ListenAndServe(":8080", nil)
}
