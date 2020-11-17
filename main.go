/*
	Бот создан в рамках проекта EBLAVOD, подробнее в телеграмм
	@ncplatzdarm @mitagmio
	или на форуме https://forum.mhub.to/
*/

package main

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/tkanos/gonfig"
)

// Configuration is 	API_KEY string  and  API_SECRET string
type Configuration struct {
	APIKey    string
	APISecret string
}

var logger *log.Logger

// Точка входа в программу
func main() {

	// В текущей директории создаем файл log.txt
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(dir+"/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	// Лог будем писать одновременно и в файл, и выводить на консоль
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)

	// При логировании будем выводить подробную информацию, включая время, название файла, строку и т.п.
	logger = log.New(mw, "", -1)
	logger.SetFlags(-1)
	logger.Println("started")

	// Открываем конфиг из json файла
	configuration := Configuration{}
	err = gonfig.GetConf("config.json", &configuration)
	if err != nil {
		panic(err)

	}

	logger.Println("api: ", configuration.APIKey, " ", configuration.APISecret)

}
