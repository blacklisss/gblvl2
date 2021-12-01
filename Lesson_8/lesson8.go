// Use ./lesson8 -dirPath=<путь к директории> [-delFiles=true]
// Аргумент dirPath не может быть пустым
package main

import (
	"context"
	"fmt"
	"gb/lvl2/Lesson_8/dto"
	"gb/lvl2/Lesson_8/filer"
	"gb/lvl2/Lesson_8/info_printer"
	"github.com/namsral/flag"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	dirPath  = flag.String("dirPath", "", "Путь к директории")
	delFiles = flag.Bool("delFiles", false, "Удалять файлы (true/false)?")

	workers = make(chan struct{}, 1000)
	errCh   = make(chan error, 1)
	doneCh  = make(chan struct{})
	wg      sync.WaitGroup
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	if err := filepath.Walk(*dirPath, filer.VisitDir); err != nil {
		log.Fatalln(err.Error())
	}

	go func() {
		for path, fi := range dto.VisitedPath {
			wg.Add(1)
			workers <- struct{}{}
			go func(path string, f os.FileInfo) {
				defer func() {
					<-workers
					wg.Done()
				}()
				err := filer.ProcessFiles(ctx, path)
				if err != nil {
					errCh <- err
					cancel()
				}
			}(path, fi)
		}
		wg.Wait()
		doneCh <- struct{}{}
	}()

L:
	for {
		select {
		case <-doneCh:
			break L
		case err := <-errCh:
			log.Fatalln(err)
		}
	}

	if *delFiles {
		if err := filer.DeleteDuplicateFiles(dto.VisitedFiles); err != nil {
			log.Fatalln(err)
		}
	} else {
		info_printer.PrintDuplicateFiles(dto.VisitedFiles)
	}
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Use %s -dirPath=<directory path> [-delFiles=true]:\n", os.Args[0])

		flag.PrintDefaults()

		fmt.Fprintf(os.Stderr, "Аргумент dirPath не может быть пустым \n")
	}

	flag.Parse()

	if *dirPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	var err error

	*dirPath, err = filepath.Abs(*dirPath)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Start...")
	info_printer.PrintLine()
}
