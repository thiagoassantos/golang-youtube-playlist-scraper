# Setup

## Authentication on YouTube API

Download your Service Account Token (JSON file) and set it in environment variable `GOOGLE_APPLICATION_CREDENTIALS`.

PowerShell example:

`$env:GOOGLE_APPLICATION_CREDENTIALS="C:\folder\file-youtube-api-token.json"`

## Run script

`go run .\main.go <YOUTUBE-PLAYLIST-ID>`

## SQLite (if you really want)

### SQLite library
* https://github.com/mattn/go-sqlite3

### Download GCC to compile

* https://jmeubank.github.io/tdm-gcc/
 
Enable the environment variable:

`$env:CGO_ENABLED=1`
