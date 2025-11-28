package utils

import (
    "io"
    "os"
)

func MergeFiles(files []string, out string) error {
    result, err := os.Create(out)
    if err != nil {
        return err
    }
    defer result.Close()

    for _, f := range files {
        part, err := os.Open(f)
        if err != nil {
            return err
        }
        _, err = io.Copy(result, part)
        part.Close()
        if err != nil {
            return err
        }
        os.Remove(f)
    }

    return nil
}
