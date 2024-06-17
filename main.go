package main

import (
	"log"
	"net/http"

	"go.etcd.io/bbolt"
)

func main() {
    db, err := bbolt.Open("bitespeed.db", 0666, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()


    if err := migrateDatabase(db); err != nil {
        log.Fatal("failed to migrate database:", err)
    }

    http.HandleFunc("/identify", identifyHandler(db))
    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
