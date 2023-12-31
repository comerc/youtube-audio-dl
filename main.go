package main

import (
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

func main() {
	videoID := "dQw4w9WgXcQ"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	// youtube Itag for m4a audio
	format := video.Formats.FindByItag(140)
	stream, _, err := client.GetStream(video, format)
	// formats := video.Formats.WithAudioChannels() // only get videos with audio
	// stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	file, err := os.Create("audio.m4a")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}
