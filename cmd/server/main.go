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

	mux.HandleFunc("POST /api/fmt/json", handleMinfyCSS)
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

	func handleMinfyCSS(w http.ResponseWriter, r *http.Request) {
		processString(w, r, formatter.MinifyCSS)
	}

	func handleFormatJson(w http.ResponseWriter, r *http.Request) {
		processString(w, r, formatter.MarshalInterface)
	}

	func handleKVtoJson(w http.ResponseWriter, r *http.Request) {
		processString(w, r, formatter.KvJson)
	}

	func handleEncode64(w http.ResponseWriter, r *http.Request) {
		processString(w, r, func(s string) (string, error)) {
			return formatter.Encodebase64(s)
		}
	}

	func handleURLEncode(w http.ResponseWriter, r *http.Request) {
		processString(w, r, func(s string) (string, error)) {
			return formatter.Encodeurl(s)
		}
	}

	func handleDecode64(w http.ResponseWriter,  r *http.Request) {
		processString(w, r, formatter.Decodebase64)
	}

	func handleURLDecode(w http.ResponseWriter,  r *http.Request) {
		processString(w, r, formatter.Decodeurl)
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

		w.header().Set("content-type", "application-json")
		json.NewEncoder(w).Encode(map[string]string{"result: pass"})
	}

	func handleUUID(w http.ResponseWriter, r *http.Request) {
		idbytes, err := generator.GenerateUUID
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		result := string(idbytes)

		w.header().Set("content-type", "application-json")
		json.NewEncoder(w).Encode(map[string]string{"result": result})
	}

	func handleSHA256(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Text string `json:"text"`
		}
		json.NewDecoder(r.Body).Decode(&req)

		result, err := generator.GenerateSHA256(req.text)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		w.header().Set("content-type", "application-json")
		json.NewEncoder(w).Encode(map[string]string{"result": result})
	}

	func handleTextStats(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Text string `json:"text"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid", 400)
		}

		chars, words, lines, nospaces, err := textutils.GetTextStats(req.text)
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

		w.header().Set("content-type", "application-json")
		json.NewEncoder(w).Encode(response)
	}

	func handleCaseConvert(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Text string `json:"text"`
			Type string `json:"type"`
		}
		json.NewEncoder(r.Body).Decode(&req)

		result := textutils.ConvertCase(eq.Text, req.Type)

		w.header().Set("content-type", "application-json")
		json.NewEncoder(w).Encode(map[string]string{"result": result})
	}
}