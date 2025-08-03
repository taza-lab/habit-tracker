package main

import (
    "backend/internal/router"
)

func main() {
    r := router.NewRouter()
    r.Run(":8080")
}
