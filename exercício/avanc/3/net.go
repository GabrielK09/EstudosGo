package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"success": false,
				"message": http.StatusInternalServerError,
			})

			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"success":  true,
			"messsage": "Hello!",
		})
	})

	http.HandleFunc("/soma", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		a := r.URL.Query()["a"]
		b := r.URL.Query()["b"]

		n1, err := strconv.Atoi(a[0])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"success": false,
				"message": err,
			})

			return
		}

		n2, err := strconv.Atoi(b[0])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"success": false,
				"message": err,
			})

			return
		}

		res := n1 + n2

		json.NewEncoder(w).Encode(map[string]any{
			"success":  true,
			"messsage": "Hello!",
			"n1":       n1,
			"n2":       n2,
			"soma":     res,
		})
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)

	}

	fmt.Println("Servidor rodando em: localhost:3000")
}
