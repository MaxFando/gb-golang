package resources

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type App struct {
	Glossary struct {
		Title    string `json:"title"`
		GlossDiv struct {
			Title     string `json:"title"`
			GlossList struct {
				GlossEntry struct {
					ID        string `json:"ID"`
					SortAs    string `json:"SortAs"`
					GlossTerm string `json:"GlossTerm"`
					Acronym   string `json:"Acronym"`
					Abbrev    string `json:"Abbrev"`
					GlossDef  struct {
						Para         string   `json:"para"`
						GlossSeeAlso []string `json:"GlossSeeAlso"`
					} `json:"GlossDef"`
					GlossSee string `json:"GlossSee"`
				} `json:"GlossEntry"`
			} `json:"GlossList"`
		} `json:"GlossDiv"`
	} `json:"glossary"`
}

func JsonDecode(filename string) {
	// Создаем файловый дескриптор для файла, в котором хрнаится json-конфигурация
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

	// Задаем переменную, в которую будем считывать конфиг
	app := App{}

	// Задаем декодер из файла и сразу же вызываем его
	err = json.NewDecoder(file).Decode(&app)
	if err != nil {
		log.Printf("Не могу декодировать json-файл в структуру: %v", err)
	}

	// Теперь в app у нас доступны параметры из файла
	fmt.Println(app.Glossary.Title)
}
