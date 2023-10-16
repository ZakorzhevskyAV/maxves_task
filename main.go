package main

import (
	"encoding/csv"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"log"
	"maxves_task/routes"
	"maxves_task/types"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	var file *os.File
	var err error

	log.Printf("opening file \"ueba.csv\"")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			log.Printf("trying to open file \"ueba.csv\"")
			file, err = os.Open("ueba.csv")
			if err != nil {
				log.Printf("failed to open file \"ueba.csv\", err: %s, trying again", err)
			} else {
				break
			}
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()
	defer file.Close()

	log.Printf("file \"ueba.csv\" opened")

	types.CSVReader = csv.NewReader(file)
	types.Records, err = types.CSVReader.ReadAll()
	if err != nil {
		log.Fatalf("failed to return the csv records, err: %s", err)
	}

	log.Printf("csv reader returned from file")

	r := mux.NewRouter()
	r.Path("/get-items/{id:[0-9]+}").HandlerFunc(routes.Get)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
