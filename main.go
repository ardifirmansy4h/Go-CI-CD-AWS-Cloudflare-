package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Halo dari Go CI/CD dari ardi !!!")
		 fmt.Fprintln(w, "Setelah ini nampil maka ci cd nya berhasil")
    })

    fmt.Println("Server running on :8080")
    http.ListenAndServe(":8080", nil)
}
