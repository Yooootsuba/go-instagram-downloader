package main

import (
    "fmt"
    "goinstagram/instagram"
)

func main() {
    fmt.Print("Enter an Instagram url : ")
    var url string
    fmt.Scanln(&url)
    downloader := instagram.NewInstagramMediaDownloader(url)
    downloader.Download()
}
