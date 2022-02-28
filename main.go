package main

import (
	"bufio"
	"fmt"
	"module34/calc"
	"module34/parserutil"
	"os"
	"strings"
)

func main() {
	outputFilename := "./output.txt"

	_ = os.Remove(outputFilename) // удаляем файл на диске
	// игнорируем ошибку т.к. файл может не существовать при первом запуске
	// а это может спровоцировать ошибку

	// используем пакет os для работы с файлами
	// указываем специальные флаги
	// O_RDWR - открыть для чтения и записи
	// O_APPEND - режим добавления текст в конец файла
	// O_CREATE - необходимо создать файл
	fout, err := os.OpenFile(outputFilename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	inputFilename := "./input.txt"

	fin, err := os.OpenFile(inputFilename, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	fileReader := bufio.NewReader(fin)

	for {
		line, _, err := fileReader.ReadLine()
		if err != nil {
			break
		}
		fmt.Println("NEW LINE:", string(line))

		str := strings.TrimSpace(string(line))
		firstNumber, secondNumber, operand, err := parserutil.ParseString(str)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			continue
		}
		calculation := calc.NewCalculator()
		result, err := calculation.Calculate(firstNumber, secondNumber, operand)
		if err != nil {
			fmt.Printf("%d%s%d error: %s\n", firstNumber, operand, secondNumber, err.Error())
			continue
		}

		fmt.Println(result)

		_, _ = fout.WriteString(fmt.Sprintf("%d%s%d=%d\n", firstNumber, operand, secondNumber, result))
	}

}
