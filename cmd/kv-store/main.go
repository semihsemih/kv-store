package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/semihsemih/kv-store/internal/router"
	"github.com/semihsemih/kv-store/internal/storage"
)

func main() {
	// Initialize a new storage
	storage := storage.New()

	// Check data.json file and if file is exist load exist data to storage
	if _, err := os.Stat("./store/data.json"); err == nil {
		dataFile, err := os.Open("./store/data.json")
		if err != nil {
			log.Println(err)
		}
		byteValue, _ := ioutil.ReadAll(dataFile)
		err = json.Unmarshal(byteValue, &storage)
		if err != nil {
			log.Println(err)
		}
		dataFile.Close()
	}

	// More control over the server's behavior is available by creating a custom Server
	srv := &http.Server{
		MaxHeaderBytes: 10,
		Addr:           ":8080",
		WriteTimeout:   time.Second * 10,
		ReadTimeout:    time.Second * 5,
		IdleTimeout:    time.Second * 60,
		Handler:        router.New(storage),
	}

	// Save to file at a specified interval (every 5 seconds)
	ticker := time.NewTicker(5 * time.Second)
	wait := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				file, _ := json.MarshalIndent(storage, "", "	")
				_ = ioutil.WriteFile("./store/data.json", file, 0644)
			case <-wait:
				ticker.Stop()
				return
			}
		}
	}()

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
