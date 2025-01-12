package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Запись в файл
func WriteFile(content interface{}, fileName string) {
	directoryPath := "C:/Users/POTishchenko/GolandProjects/testComposition"
	fullPath := filepath.Join(directoryPath, fileName)

	// Создание директории
	if err := os.MkdirAll(directoryPath, os.ModePerm); err != nil {
		fmt.Println("Ошибка при создании директории: ", err)
		return
	}
	// Создание файла
	file, err := os.Create(fullPath)
	if err != nil {
		fmt.Println("Ошибка при создании файла: ", err)
		return
	}
	defer file.Close()

	// Запись файла
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ") // устанавливаем отступ

	if err := encoder.Encode(content); err != nil {
		fmt.Println("Ошибка при записи в файл: ", err)
		return
	}
	fmt.Println("Файл создан и записан")
}

// Чтение файла
func ReadFile(fileName string) (content string, err error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Ошибка при чтении файла: ", err)
		return "", err
	}
	return string(data), nil
}
