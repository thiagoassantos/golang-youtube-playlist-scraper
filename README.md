# Setup

## Authentication on YouTube API

Download your Service Account Token (JSON file) and set it in environment variable:

`$env:GOOGLE_APPLICATION_CREDENTIALS="C:\folder\file-youtube-api-token.json"`

## SQLite (if you really want)

### Download GCC to compile

* https://jmeubank.github.io/tdm-gcc/
 
### Enable the environment variable
`$env:CGO_ENABLED=1`

### SQLite library
* https://github.com/mattn/go-sqlite3