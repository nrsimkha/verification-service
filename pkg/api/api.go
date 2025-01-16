package api

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"verification/pkg/logger"

	"github.com/gorilla/mux"
)

type API struct {
	r *mux.Router
}

type Comment struct {
	ID       int
	TextBody string
	PubTime  int64
	NewsID   int
	ParentID sql.NullInt64
}

var spamWords = []string{"qwerty", "йцукен", "zxvbnm"}

// Конструтор API
func New() *API {
	api := API{}
	api.r = mux.NewRouter()
	api.endpoints()
	api.r.Use(logger.WrapHandlerWithLogging)
	return &api
}

// Router возвращает маршрутизатор запросов.
func (api *API) Router() *mux.Router {
	return api.r
}

func (api *API) endpoints() {
	api.r.HandleFunc("/verification", checkTextHandler).Methods(http.MethodPost)

}

func checkTextHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read response body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var c Comment
	err = json.Unmarshal(body, &c)
	if err != nil {
		log.Printf("failed to marshal response: %v", err)
		http.Error(w, "invalid body structure", http.StatusBadRequest)
		return
	}
	if isSpam(c) {
		w.WriteHeader(400)
	} else {
		w.WriteHeader(200)
	}

}

func isSpam(c Comment) bool {
	for _, spamWord := range spamWords {
		if strings.Contains(c.TextBody, spamWord) {
			return true
		}
	}
	return false
}
