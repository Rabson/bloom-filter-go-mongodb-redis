
package main

import (
    "log"
    "net/http"
    "username-check-api/internal/api"
    "username-check-api/internal/bloom"
    "username-check-api/internal/db"
)

func main() {
    bloom.InitBloom()
    db.ConnectMongo()

    router := api.SetupRoutes()
    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
