// Package dto предоставляеет структуры данных для программы поиска дубликатов файла
package dto

import "os"

type FileInfo struct {
	Name string
	Path string
}

type FilesInfo struct {
	Files []FileInfo
	Count uint64
}

type FilesMap map[string]FilesInfo

var (
	VisitedPath  map[string]os.FileInfo
	VisitedFiles FilesMap
)

func init() {
	VisitedPath = make(map[string]os.FileInfo)
	VisitedFiles = make(FilesMap)
}
