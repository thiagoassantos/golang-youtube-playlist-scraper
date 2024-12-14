package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/youtube/v3"
)

// List all videos from a given playlist
func lsPlaylistVideos(yts *youtube.Service, pid string) error {
	q := yts.PlaylistItems.List([]string{"snippet"})

	var pt = ""

	for {
		xs, err := q.Do(
			googleapi.QueryParameter("playlistId", pid),
			googleapi.QueryParameter("maxResults", "7"),
			googleapi.QueryParameter("pageToken", pt),
		)
		if err != nil {
			return err
		}

		for _, x := range xs.Items {

			fmt.Println(
				lsVideoDetails(yts, x.Snippet.ResourceId.VideoId),
			)
		}

		if err != nil {
			log.Fatal(err)
		}

		pt = xs.NextPageToken

		if pt == "" {
			return nil
		}

		fmt.Println(" ----- ")
	}
}

// Get video details for database
func lsVideoDetails(yts *youtube.Service, vid string) error {

	xs, err := yts.Videos.List([]string{"snippet"}).Do(
		googleapi.QueryParameter("id", vid),
	)

	if err != nil {
		return err
	}

	for _, x := range xs.Items {

		db, err := sql.Open("sqlite3", "./database.db")
		checkErr(err)
		// insert
		stmt, err := db.Prepare("INSERT INTO videoinfo(videoId, channelArtist, songTitle, url, thumbnail, collected) values(?,?,?,?,?,?)")
		checkErr(err)

		// Some channel names come with the suffix " - Topic", so it is filtered.
		res, err := stmt.Exec(vid, strings.TrimRight(x.Snippet.ChannelTitle, " - Topic"), x.Snippet.Title, "https://www.youtube.com/watch?v="+vid, x.Snippet.Thumbnails.High.Url, 1)
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)

		fmt.Println(id)

		db.Close()
	}

	return nil
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please, insert a YouTube playlist ID.")
		return
	}

	playlistLink := os.Args[1]
	fmt.Println("Argumento: ", playlistLink)

	yts, err := youtube.NewService(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = lsPlaylistVideos(yts, playlistLink)

	db, err := sql.Open("sqlite3", "./database.db")
	checkErr(err)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM videoinfo")
	checkErr(err)
	var uid string
	var channelArtist string
	var songTitle string
	var url string
	var thumbnail string
	var collected string

	for rows.Next() {
		err = rows.Scan(&uid, &channelArtist, &songTitle, &url, &thumbnail, &collected)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(channelArtist)
		fmt.Println(songTitle)
		fmt.Println(url)
		fmt.Println(thumbnail)
		fmt.Println(collected)
	}

	rows.Close()

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
