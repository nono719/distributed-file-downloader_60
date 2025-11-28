package main

import (
    "flag"
    "fmt"
    "distributed-file-downloader_60/internal/downloader"
)

func main() {
    url := flag.String("u", "", "Download URL")
    out := flag.String("o", "output.file", "Output filename")
    workers := flag.Int("t", 8, "Worker threads")

    flag.Parse()

    if *url == "" {
        fmt.Println("Usage: downloader -u <url> -o <output> -t <threads>")
        return
    }

    d := downloader.New(*url, *out, *workers)
    err := d.Start()
    if err != nil {
        fmt.Println("Download failed:", err)
    } else {
        fmt.Println("Download completed:", *out)
    }
}
