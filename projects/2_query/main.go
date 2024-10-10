package main

import (
	"fmt"

	"net/http" // пакет для поддержки HTTP протокола
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello," + r.URL.Query().Get("name") + "!"))
	if err != nil {
		panic(err)
	}

}
func main() {
	// здесь ваш код
	http.HandleFunc("/api/user", handler)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера: ", err)
	}
}
