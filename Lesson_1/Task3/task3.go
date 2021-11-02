package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	path := "../tmp"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModeDir)
	}

	for i := 0; i < 1000000; i++ {
		filename := path + "/file_" + strconv.Itoa(i) + ".txt"
		if err := createFileAndPushContent(filename); err != nil {
			fmt.Printf("[Error] %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("[INFO] %d из 1 000 000\n", i)
	}

	fmt.Println("[INFO] Done")
}

func createFileAndPushContent(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close() //Закрываем дескриптор файла при выходе из функции

	//Получаем и записываем какой-то контент

	return nil
}
