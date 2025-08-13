package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type BarcodeSimples struct {
	Code     string
	Datahora string
}

type Barcodes struct {
	Barcodes []BarcodeSimples
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func write(filename string, bs Barcodes) {
	data, err := json.MarshalIndent(bs, "", "  ")
	if err != nil {
		log.Fatal("Erro ao criar JSON:", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal("Erro ao criar arquivo:", err)
	}

	fmt.Println("Arquivo criado com sucesso:", filename)
}

func writeToFileSimples(barcode string) {
	filename := "simples.json"

	if !fileExists(filename) {
		fmt.Println("Arquivo não existe, criando com dados padrão...")

		bs := Barcodes{
			Barcodes: []BarcodeSimples{
				{Code: barcode, Datahora: time.Now().Format("02/01/2006 15:04")},
			},
		}
		write(filename, bs)
	}

	// Agora lê o arquivo
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Erro ao abrir arquivo:", err)
	}
	defer file.Close()

	var bs Barcodes
	err = json.NewDecoder(file).Decode(&bs)
	if err != nil {
		log.Fatal("Erro ao decodificar JSON:", err)
	}

	bs.Barcodes = append(bs.Barcodes, BarcodeSimples{Code: barcode, Datahora: time.Now().Format("02/01/2006 15:04")})
	write(filename, bs)

}

func resetSimples(){
	filename := "simples.json"

	if !fileExists(filename) {
		return
	}
	write(filename, Barcodes{Barcodes: []BarcodeSimples{}})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/simples/{barcode}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		barcode := vars["barcode"]
		writeToFileSimples(barcode)
	}).Methods("GET")

	r.HandleFunc("/simples/reset", func(w http.ResponseWriter, r *http.Request) {
		resetSimples()
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}
