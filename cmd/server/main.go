package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"

	"github.com/kartikangiras/text-forge/internal"
)

var uiAssets embed.FS

type StringProcessor func(string) (string, error)

func main() {
	distFS, err := fs.Sub(uiAssets, "ui/dist")
	if err != nil {
		log.Fatal(err)

		mux := http.NewServeMux()

		mux.HandleFunc("POST /api/fmt/css", makeHandler(internal.MinifyCSS))
		mux.HandleFunc("POST /api/fmt/json", makeHandler(internal.MarshalInterface))
		mux.HandleFunc("POST /api/fmt/kvjson", makeHandler(internal.KvJson))
		mux.HandleFunc("POST /api/fmt/b64dec", makeHandler(internal.Decodebase64))
		mux.HandleFunc("POST /api/fmt/b64en", makeHandler(internal.Encodebase64))
		mux.HandleFunc("POST /api/fmt/urldec", makeHandler(internal.Decodeurl))
		mux.HandleFunc("POST /api/fmt/urlen", makeHandler(internal.Encodeurl))

		mux.HandleFunc("POST /api/fmt/clean", makeHandler(internal.CleanUpText))

		mux.HandleFunc("POST /api/fmt/sha256", makeHandler(internal.GenerateSHA256))

		mux.HandleFunc("POST /api/fmt/uuid", handleUUID)
		mux.HandleFunc("POST /api/fmt/pass", handlePassword)
		mux.HandleFunc("POST /api/fmt/case", handleCaseConvert)
		mux.HandleFunc("POST /api/fmt/stats", handleTextStats)

		fileServer := http.FileServer(http.FS(distFS))
		mux.Handle("/", fileServer)

		log.Println("Server starting on http://localhost:8080")
		if err := http.ListenAndServe(":8080", mux); err != nil {
			log.Fatal(err)
		}
	}
}

func handlePassword(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Length int `json:"length"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid", 400)
	}

	pass, err := internal.GeneratePassword(req.Length)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application-json")
	json.NewEncoder(w).Encode(map[string]string{"result": pass})
}

func handleTextStats(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid", 400)
	}

	chars, words, lines, nospaces, err := internal.GetTextStats(req.Text)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	response := map[string]int{
		"characters": chars,
		"words":      words,
		"lines":      lines,
		"nospaces":   nospaces,
	}

	w.Header().Set("content-type", "application-json")
	json.NewEncoder(w).Encode(response)
}

func handleCaseConvert(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Text string `json:"text"`
		Type string `json:"type"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid", 400)
		return
	}

	result := internal.ConvertCase(req.Text, req.Type)

	w.Header().Set("content-type", "application-json")
	json.NewEncoder(w).Encode(map[string]string{"result": string(result)})
}

func makeHandler(processor StringProcessor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Text string `json:"text"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid", 400)
			return
		}

		result, err := processor(req.Text)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"result": result})
	}
}

func handleUUID(w http.ResponseWriter, r *http.Request) {
	result, err := internal.GenerateUUID()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": string(result)})
}
