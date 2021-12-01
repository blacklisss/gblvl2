package filer

import (
	"context"
	"gb/lvl2/Lesson_8/dto"
	"os"
	"testing"
)

func TestDelFileError(t *testing.T) {

	path := "/dev/null"

	if err := DelFile(path); err == nil {
		t.Error("ошибка удаления. ожидалась ошибка.")
	}
}

func TestDelFileSuccess(t *testing.T) {
	var data = []struct {
		name string
		perm os.FileMode
	}{
		{"test1.file", 0666},
		{"test2.file", 0666},
		{"test3.file", 0666},
	}

	// Создаем тестовые файлы
	for _, fData := range data {
		f, err := os.Create(fData.name)
		if err != nil {
			f.Close()
			panic(err)
		}

		f.Close()
	}

	for _, fData := range data {
		if err := DelFile(fData.name); err != nil {
			t.Error(err)
		}
	}
}

func TestGetCRC64SumString(t *testing.T) {
	path := "filer.go"

	f, err := os.Open(path)
	if err != nil {
		t.Error("error open file:", path)
	}
	defer f.Close()

	f.Seek(0, 0)
	_, err = GetCRC64SumString(f)
	if err != nil {
		t.Error("error:", err)
	}
}

func TestProcessFilesError(t *testing.T) {
	path := "/dev/null"
	ctx, _ := context.WithCancel(context.Background())

	if err := ProcessFiles(ctx, path); err == nil {
		t.Error("ожидалась ошибка")
	}
}

func TestProcessFilesOpenError(t *testing.T) {
	path := "/dfdf/"
	ctx, _ := context.WithCancel(context.Background())

	fileInfo, err := os.Lstat("filer.go")
	if err != nil {
		t.Error(err)
		os.Exit(1)
	}

	dto.VisitedPath = make(map[string]os.FileInfo)
	dto.VisitedPath[path] = fileInfo

	if err = ProcessFiles(ctx, path); err == nil {
		t.Error("ожидалась ошибка")
	}
}

func TestProcessFilesCheckExistFile(t *testing.T) {
	path := "filer.go"
	ctx, _ := context.WithCancel(context.Background())

	fileInfo, err := os.Lstat(path)
	if err != nil {
		t.Error(err)
		os.Exit(1)
	}

	dto.VisitedPath = make(map[string]os.FileInfo)
	dto.VisitedPath[path] = fileInfo

	dto.VisitedFiles = make(dto.FilesMap)

	f, err := os.Open(path)
	if err != nil {
		t.Error("error open file:", path)
	}
	defer f.Close()

	f.Seek(0, 0)
	sum, err := GetCRC64SumString(f)
	if err != nil {
		t.Error("error:", err)
	}

	ftmp := []dto.FileInfo{
		{
			path,
			path,
		},
	}
	dto.VisitedFiles[sum] = dto.FilesInfo{
		Files: ftmp,
		Count: 1,
	}

	if err = ProcessFiles(ctx, path); err != nil {
		t.Error(err)
	}
}
func TestProcessFilesCheckNonExistFile(t *testing.T) {
	path := "filer.go"
	ctx, _ := context.WithCancel(context.Background())

	fileInfo, err := os.Lstat(path)
	if err != nil {
		t.Error(err)
		os.Exit(1)
	}

	dto.VisitedPath = make(map[string]os.FileInfo)
	dto.VisitedPath[path] = fileInfo

	dto.VisitedFiles = make(dto.FilesMap)

	f, err := os.Open(path)
	if err != nil {
		t.Error("error open file:", path)
	}
	defer f.Close()

	f.Seek(0, 0)
	sum, err := GetCRC64SumString(f)
	if err != nil {
		t.Error("error:", err)
	}

	if err = ProcessFiles(ctx, path); err != nil {
		t.Error(err)
	}

	if _, ok := dto.VisitedFiles[sum]; !ok {
		t.Error("ощибка заполнения картыы посещенных файлов")
	}
}

func TestVisitDir(t *testing.T) {
	fileInfo, err := os.Lstat("filer.go")
	if err != nil {
		t.Error(err)
		os.Exit(1)
	}

	err = VisitDir(".", fileInfo, err)
	if err != nil {
		t.Error(err)
	}

}

func TestVisitDirDir(t *testing.T) {

	path := "test"
	err := os.Mkdir(path, os.ModeDir)
	if err != nil {
		t.Error(err)
		os.Exit(1)
	}

	fileInfo, err := os.Lstat(path)
	if err != nil {
		t.Error(err)
		os.Exit(1)
	}

	defer os.Remove(path)

	err = VisitDir(path, fileInfo, err)
	if err != nil {
		t.Error(err)
	}

}

func TestVisitDirCheckPath(t *testing.T) {
	path := "filer.go"
	fileInfo, err := os.Lstat(path)
	if err != nil {
		t.Error(err)
		os.Exit(1)
	}

	dto.VisitedPath = make(map[string]os.FileInfo)
	dto.VisitedPath[path] = fileInfo

	err = VisitDir(path, fileInfo, err)
	if err != nil {
		t.Error(err)
	}

}
