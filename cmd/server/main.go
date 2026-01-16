package main

import (
	"embed"
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

	mux.HandleFunc("POST /api/fmt/css", makeHandler(formatter.MinifyCSS))
	mux.HandleFunc("POST /api/fmt/json", makeHandler(formatter.formatter.MarshalInterface))
	mux.HandleFunc("POST /api/fmt/kvjson", makeHandler(formatter.KvJson))
	mux.HandleFunc("POST /api/fmt/b64dec", makeHandler(formatter.Decodebase64))
	mux.HandleFunc("POST /api/fmt/b64en", makeHandler(formatter.Encodebase64))
	mux.HandleFunc("POST /api/fmt/urldec", makeHandler(formatter.Decodeurl))
	mux.HandleFunc("POST /api/fmt/urlen", makeHandler(formatter.Encodeurl))

	mux.HandleFunc("POST /api/fmt/clean", makeHandler(textutils.CleanUpText))

	mux.HandleFunc("POST /api/fmt/sha256", makeHandler(generator.GenerateSHA256))

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

		pass, err := generator.GeneratePassword(req.Length)
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

		chars, words, lines, nospaces, err := textutils.GetTextStats(req.Text)
		if err != nil {
			http.Error(w, err.Error(), 400);
			return
		}

		response := map[string]int {
			"characters": chars,
			"words": words,
			"lines": lines,
			"nospaces": nospaces,
		}

		w.Header().Set("content-type", "application-json")
		json.NewEncoder(w).Encode(response)
	}

	func handleCaseConvert(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Text string `json:"text"`
			Type string `json:"type"`
		}
		json.NewEncoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid", 400)
			return
		}

		result := textutils.ConvertCase(req.Text, req.Type)

		w.Header().Set("content-type", "application-json")
		json.NewEncoder(w).Encode(map[string]string{"result": result})
	}

	func makeHandler(processor StringProcessor) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var req struct {
				Text string `json:"text"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "invalid", 400)
			}

			result, err := processor(req.Text)
			if err != nil {
				http.Error(w, err.Error(), 400);
				return
			}

			w.header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"result": result})
		}
	}

	func handleUUID(w http.ResponseWriter, r *http.Request) {
	result := generator.GenerateUUID()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": result})
	}
