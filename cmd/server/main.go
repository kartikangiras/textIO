package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/kartikangiras/text-forge/internal"
)

var uiAssets embed.FS

func main() {
	distFS, err := fs.Sub(uiAssets, "ui/dist")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/fmt/json", handleMinfiyCSS)
	mux.HandleFunc("POST /api/fmt/json", handleFormatJson)
	mux.HandleFunc("POST /api/fmt/json", handleKVtoJson)
	mux.HandleFunc("POST /api/fmt/json", handleEncode64)
	mux.HandleFunc("POST /api/fmt/json", handleDecode64)
	mux.HandleFunc("POST /api/fmt/json", handleURLEncode)
	mux.HandleFunc("POST /api/fmt/json", handleURLDecode)

	mux.HandleFunc("POST /api/fmt/json", handlePassword)
	mux.HandleFunc("POST /api/fmt/json", handleSHA256)
	mux.HandleFunc("GET /api/fmt/json", handleUUID)

	mux.HandleFunc("POST /api/fmt/json", handleCleanUp)
	mux.HandleFunc("POST /api/fmt/json", handleCaseConvert)
	mux.HandleFunc("POST /api/fmt/json", handleTextStats)

	fileServer := http.FileServer(http.FS(distFS))
	mux.Handle("/", fileServer)

	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
