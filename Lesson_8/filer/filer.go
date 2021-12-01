// Package filer используется для обхода заданной директории, сохранения имен файлов и информации о них в структуры
// dto.VisitedPath и dto.VisitedFiles
package filer

import (
	"context"
	"errors"
	"fmt"
	"gb/lvl2/Lesson_8/dto"
	"gb/lvl2/Lesson_8/info_printer"
	"hash/crc64"
	"io"
	"os"
	"strconv"
	"sync"
)

var (
	crc64q = crc64.MakeTable(0xC96C5795D7870F42)
	lock   sync.Mutex
)

// VisitDir используется для рекурсивного обхода директории
// заполняет структуру dto.VisitedPath
// возвращает ошибку или nil
func VisitDir(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		return nil
	}

	if _, ok := dto.VisitedPath[path]; ok {
		return nil
	}

	dto.VisitedPath[path] = f

	return nil
}

// ProcessFiles обрабатывает файлы, находит дубликаты с помощью хеша CRC64
// заполняет структуру dto.VisitedFiles
// возвращает ошибку или nil
func ProcessFiles(ctx context.Context, path string) error {
	if _, ok := dto.VisitedPath[path]; !ok {
		return errors.New("нет такого пути:" + path)
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	f := dto.VisitedPath[path]

	file.Seek(0, 0)
	sum, err := GetCRC64SumString(file)

	lock.Lock()
	if fileinfo, ok := dto.VisitedFiles[sum]; !ok {
		fileinfo = dto.FilesInfo{}
		fileinfo.Files = append(fileinfo.Files, dto.FileInfo{
			f.Name(),
			path,
		})
		fileinfo.Count = 1

		dto.VisitedFiles[sum] = fileinfo
	} else {
		fileinfo.Files = append(fileinfo.Files, dto.FileInfo{
			f.Name(),
			path,
		})
		fileinfo.Count += 1

		dto.VisitedFiles[sum] = fileinfo
	}
	lock.Unlock()
	return nil
}

// GetCRC64SumString возвращает хэш CRC64 содержимого файла
func GetCRC64SumString(f *os.File) (string, error) {
	file1Sum := crc64.New(crc64q)
	_, err := io.Copy(file1Sum, f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", file1Sum.Sum(nil)), nil
}

// DeleteDuplicateFiles удаляет выбранные пользователем дубликаты файлов
func DeleteDuplicateFiles(files map[string]dto.FilesInfo) error {
	var action string
	for _, vf := range files {
		if vf.Count > 1 {
			fmt.Println("Выберите файл для удаления:")
			var i uint64
			for i = 0; i < vf.Count; i++ {
				fmt.Printf("%d) %s : %s\n", i+1, vf.Files[i].Name, vf.Files[i].Path)
			}

			fmt.Printf("all – Все файлы\n")
			fmt.Printf("skip – Пропустить удаление\n")

			for {
				if _, err := fmt.Scanln(&action); err != nil {
					return err
				}

				if action == "skip" {
					fmt.Println("Пропускаем удаление")
					info_printer.PrintLine()
					break
				} else if action == "all" {
					var j uint64
					for j = 0; j < vf.Count; j++ {
						fmt.Println("Удаляем файл", vf.Files[j])
						if err := DelFile(vf.Files[j].Path); err != nil {
							fmt.Printf("ошибка удаления файла %s: %v", vf.Files[j].Path, err)
						}
					}
					info_printer.PrintLine()
					break
				} else {
					if num, err := strconv.Atoi(action); err != nil {
						fmt.Println("Необходимо ввести либо число, либо ключевые слова")
					} else if uint64(num) > vf.Count {
						fmt.Println("Необходимо ввести число от 1 до", vf.Count)
					} else {
						fmt.Println("Удаляем файл", vf.Files[num-1])
						if err = DelFile(vf.Files[num-1].Path); err != nil {
							fmt.Printf("ошибка удаления файла %s: %v", vf.Files[num-1].Path, err)
						}
						info_printer.PrintLine()
						break
					}
				}
			}
		}
	}

	return nil
}

// DelFile удаляет файл по переданному в функцию пути
// возвращает ошибку удаления или nil
func DelFile(path string) error {
	return os.Remove(path)
}
