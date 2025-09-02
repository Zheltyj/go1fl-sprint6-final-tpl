package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Zheltyj/go1fl-sprint6-final-tpl/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("Unsupported method %s", r.Method), http.StatusInternalServerError)
		return
	}
	data, err := os.ReadFile("./index.html")
	if err != nil {
		http.Error(w, "Read error index.html", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Unsupported method %s", r.Method), http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error receiving file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "File read error", http.StatusInternalServerError)
		return
	}

	convertedData := service.ConvertString(string(data))

	fileName := fmt.Sprintf("../%s%s", time.Now().UTC().Format("2006-01-02_15-04-05"), filepath.Ext("output.txt"))

	err = os.WriteFile(fileName, []byte(convertedData), 0755)
	if err != nil {
		http.Error(w, "File write error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(convertedData))
}
