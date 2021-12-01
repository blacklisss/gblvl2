// Package info_printer выводит на печать найденные дубликаты файлов и линию разграничения
package info_printer

import (
	"fmt"
	"gb/lvl2/Lesson_8/dto"
	"strings"
)

func PrintDuplicateFiles(files dto.FilesMap) {
	for _, vf := range files {
		if vf.Count > 1 {
			var i uint64
			for i = 0; i < vf.Count; i++ {
				fmt.Println(vf.Files[i].Name, ":", vf.Files[i].Path)
			}

			PrintLine()
		}
	}

}

func PrintLine() {
	fmt.Println(strings.Repeat("-", 20))
}
