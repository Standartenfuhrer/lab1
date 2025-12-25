package main

import (
    "fmt"
    "io/fs"
    "os"
    "path/filepath"
    "sort"
    "sync"
)


type FileResult struct {
    Path string
    Size int64
}
func (fr FileResult) String() string {
    const unit = 1024
    if fr.Size < unit {
        return fmt.Sprintf("%d B - %s", fr.Size, fr.Path)
    }

    div, exp := int64(unit), 0
    for n := fr.Size / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return fmt.Sprintf("%1.f%cB - %s", float64(fr.Size)/float64(div), "KMGTPE"[exp], fr.Path)
}
type Processor interface {
    Processor(path string, entry fs.DirEntry) (FileResult, error)
}

type SizeProcessor struct{}

func (sp SizeProcessor) Processor(path string, entry fs.DirEntry) (FileResult, error) {
    info, err := entry.Info()
    if err != nil {
        return FileResult{}, err
    }
    return FileResult{Path: path, Size: info.Size()}, nil
}
var sem = make(chan struct{}, 20)

func walk(path string, processor Processor, results chan<- FileResult, wg *sync.WaitGroup) {
    entries, err := os.ReadDir(path)
    if err != nil {
        return
    }

    for _, entry := range entries {
        fullPath := filepath.Join(path, entry.Name())

        if entry.IsDir() {
        } else {
            wg.Add(1)
            go func(p string, e fs.DirEntry) {
                defer wg.Done()

                sem <- struct{}{}
                defer func() { <-sem }()

                res, err := processor.Processor(p, e)
                if err == nil {
                    results <- res
                }
            }(fullPath, entry)
        }
    }
}

// 3. --- Main ---
func main() {
    root := "C:/tamgo/project2/lab1"
    if len(os.Args) > 1 {
        root = os.Args[1]
    }

    resultsChan := make(chan FileResult, 100)
    var wg sync.WaitGroup

    go func() {
        defer close(resultsChan)
        walk(root, SizeProcessor{}, resultsChan, &wg)
        wg.Wait()
    }()

    var allFiles []FileResult
    for res := range resultsChan {
        allFiles = append(allFiles, res)
    }

    sort.Slice(allFiles, func(i, j int) bool {
        return allFiles[i].Size > allFiles[j].Size
    })

    fmt.Println("--- Largest 5 files ---")
    limit := 5
    if len(allFiles) < limit {
        limit = len(allFiles)
    }
    for i := 0; i < limit; i++ {
        fmt.Println(allFiles[i])
    }
}
