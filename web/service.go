package web

import (
	"log"
	"net/http"
	"text/template"

	"github.com/SmartParksOrg/smartparks-connect-web/device_template"
	"github.com/SmartParksOrg/smartparks-connect-web/grpc_client"
)

// handleRoot handles requests to the root URL ("/").
func (h *Handler) handleRoot(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/dist/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

// handleTemplateList handles GET requests to retrieve a list of templates.
func (h *Handler) handleTemplateList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	Succ(w, h.templateManager.GetTemplates())
}

func StartHttpServer(client *grpc_client.GrpcClient, template *device_template.TemplateManager) {
	handle := Handler{
		grpcClient:      client,
		templateManager: template,
	}
	http.Handle("/api/v1/login", cors(http.HandlerFunc(handle.handleLogin)))
	http.Handle("/api/v1/device/queue", cors(http.HandlerFunc(handle.handleAPI)))
	http.Handle("/api/v1/list", cors(http.HandlerFunc(handle.handleList)))
	http.Handle("/api/v1/template/list", cors(http.HandlerFunc(handle.handleTemplateList)))

	// RockBLOCK API
	http.Handle("/api/v1/rockblock/queue", cors(http.HandlerFunc(handle.handleRockBLOCKAPI)))
	http.Handle("/api/v1/rockblock/login", cors(http.HandlerFunc(handle.handleRockBLOCKLogin)))

	http.Handle("/assets/device_template/", http.StripPrefix("/assets/device_template/", http.FileServer(http.Dir("template_files"))))
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("public/dist/assets"))))
	http.HandleFunc("/", handle.handleRoot)
	log.Println("http server runing in port : 8881")
	log.Fatal(http.ListenAndServe(":8881", nil))
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set headers to allow all CORS requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

		// If it's an OPTIONS request, return immediately with a 200 status code
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next middleware in the chain
		next.ServeHTTP(w, r)
	})
}
