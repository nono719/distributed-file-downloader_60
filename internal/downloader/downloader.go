package downloader

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strconv"
    "sync"
    "distributed-file-downloader_60/internal/utils"
)

type Downloader struct {
    URL     string
    Out     string
    Workers int
    Size    int64
}

func New(url, out string, workers int) *Downloader {
    return &Downloader{URL: url, Out: out, Workers: workers}
}

func (d *Downloader) Start() error {
    resp, err := http.Head(d.URL)
    if err != nil {
        return err
    }
    if resp.StatusCode != 200 {
        return fmt.Errorf("HTTP error: %s", resp.Status)
    }

    size, _ := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
    d.Size = size

    partSize := d.Size / int64(d.Workers)
    var wg sync.WaitGroup

    tempFiles := make([]string, d.Workers)

    for i := 0; i < d.Workers; i++ {
        wg.Add(1)
        start := int64(i) * partSize
        end := start + partSize - 1
        if i == d.Workers-1 {
            end = d.Size - 1
        }

        tmp := fmt.Sprintf("%s.part%d", d.Out, i)
        tempFiles[i] = tmp

        go func(start, end int64, idx int) {
            defer wg.Done()
            d.downloadPart(start, end, tempFiles[idx])
        }(start, end, i)
    }

    wg.Wait()
    return utils.MergeFiles(tempFiles, d.Out)
}

func (d *Downloader) downloadPart(start, end int64, filename string) error {
    req, _ := http.NewRequest("GET", d.URL, nil)
    req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    out, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, resp.Body)
    return err
}
