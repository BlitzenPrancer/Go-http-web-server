package main

import (
	"encoding/json"
	"io/ioutil"
)

// defining struct on what the video looks like
type video struct {
	Id          string
	Title       string
	Description string
	Imageurl    string
	Url         string
}

// function to retrive the list of videos from an user's file
func getVideos() (videos []video) {
	fileBytes, err := ioutil.ReadFile("./videos.json")
	// checking for error and handling
	if err != nil {
		panic(err)
	}
	// using json.Unmarshal to convert to convert slice of bytes into videos slice
	err = json.Unmarshal(fileBytes, &videos)
	if err != nil {
		panic(err)
	}
	return videos
}

// function to save the videos
func saveVideos(videos []video) {
	// converting video into slice of bytes
	videoBytes, err := json.Marshal(videos)
	// checking and handling errors
	if err != nil {
		panic(err)
	}
	// saving the slice of bytes by passing to ioutil.WriteFile function
	err = ioutil.WriteFile("./videos-updated.json", videoBytes, 0644)
	// checking and handling errors
	if err != nil {
		panic(err)
	}
}
