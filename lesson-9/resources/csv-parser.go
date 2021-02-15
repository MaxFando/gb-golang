package resources

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	maxRows = flag.Int("max-rows", 5, "Максимальное число строк, которые нужно прочитать")
	columns = flag.String("columns", "", "Разделенный запятыми список столбцов, которые нужно вывести")
)

func main() {
	flag.Parse() // Сообщаем библиотеке flag о том, что необходимо считать флаги

	// Имя файла передается не через флаги, а как аргументы командной строки flag.Arg.
	if flag.NArg() != 1 {
		log.Fatal(
			"Неверно задано количество аргументов командной строки. Проверьте, что вы задали имя файла.",
		)
	}

	filename := strings.TrimSpace(flag.Arg(0)) // Считываем имя файла и очищаем ввод от пробелов

	// Проверяем, что файл существует
	if _, err := os.Stat(filename); err != nil {
		log.Fatalf("Не могу проверить существование файла: %v", err)
	}

	// Пробуем открыть файл
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Не могу открыть файл: %v", err)
	}

	// Не забываем закрыть файл при выходе из функции
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Не могу закрыть файл: %v", err)
		}
	}()

	// Проверяем флаги
	if *maxRows < 1 || *maxRows > 100 {
		log.Fatalf("Значение max-rows должно быть от 1 до 100")
	}

	// Разбиваем заданные колонки по столбцам
	if len(*columns) == 0 {
		log.Fatal("Не заданы столбцы для чтения из файла")
	}
	columns := strings.Split(*columns, ",")

	// Проверяем наличие колонок в файле:
	// Создаем мапу, в которую будем писать номера колонок с заданным именем,
	// считываем первую строку CSV-файла и раздаем номера колонкам
	colNumbers := make(map[string]int, len(columns))
	reader := csv.NewReader(file)  // Считываем файл с помощью библиотеки encoding/csv
	fileCols, err := reader.Read() //  Считываем заголовочную строку
	if err != nil {
		log.Fatalf("Не могу считать заголовочную строку csv-файла: %v", err)
	}

	for i, col := range fileCols {
		for j, c := range columns {
			if col == c { // Если колонка есть в списке желаемых
				colNumbers[c] = i                               // Добавляем индекс колонки в мапу
				columns = append(columns[:j], columns[j+1:]...) // "Помечаем" колонку обработанной - удаляем ее из спика
			}
		}
	}

	// Если после прохождения по всем колонкам файла в исходном списке желаемых колонок
	// что-то осталось, значит, эта колонка не существует в файле - выводим ошибку
	if len(columns) > 0 {
		log.Fatalf("В файле не найдены столбцы: %v", columns)
	}

	// Читаем заданное кол-во строк
	for i := 1; i <= *maxRows; i++ {
		row, err := reader.Read()

		// Если в качестве ошибки получили EOF (End of file), значит
		// файл закончился раньше, чем мы с читали максимально заданное кол-во строк.
		// прекращаем чтение
		if err == io.EOF {
			break
		}

		// Если произошла какая-то другая ошибка, прекращаем работу программы
		if err != nil {
			log.Fatalf("Ошибка при чтении файла: %v", err)
		}

		// Выводим значения строки в ожидаемом формате:
		fmt.Printf("---------- Строка %d ---------\n", i)
		for key, col := range colNumbers {
			fmt.Printf("%s:\t%s\n", key, row[col])
		}
		fmt.Printf("-----------------------------\n\n")
	}
}
