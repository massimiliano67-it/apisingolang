package main

import (
	"apimusic/pkg/music"
	"apimusic/pkg/sampledata"
	"apimusic/pkg/server"
	"apimusic/pkg/storage/simdata"
	"flag"
	"fmt"
	"log"
	"net/http"
)

// main starting program
func main() {

	datasim := flag.Bool("datasim", false, "initialize the api with some music")
	flag.Parse()

	var musics map[string]*music.Music
	//musics = sampledata.Music

	if *datasim {
		fmt.Println("DATASIM OK")
		musics = sampledata.Music
	} else {
		fmt.Println("DATASIM NO", musics)
	}
	musics = sampledata.Music
	repo := simdata.NewMusicRepository(musics)
	s := server.New(repo)

	fmt.Println("The music server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
