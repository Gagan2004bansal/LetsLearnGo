package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Gagan2004bansal/LetsLearnGo/internal/model"
	"github.com/Gagan2004bansal/LetsLearnGo/internal/service"
	"github.com/gorilla/mux"
)

type UrlHandler struct {
	urlService *service.UrlService
}

func NewUrlHandler(urlService *service.UrlService) *UrlHandler {
	return &UrlHandler{urlService: urlService}
}

func (h *UrlHandler) UrlRoutes(router *mux.Router) {
	router.HandleFunc("/url", h.CreateUrl).Methods(http.MethodPost)
}

func (h *UrlHandler) CreateUrl(w http.ResponseWriter, r *http.Request) {
	var request model.ReqUrl

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
	}

	if request.Url == "" {
		http.Error(w, "please add url", http.StatusBadRequest)
		return
	}

	slog.Info("Everything okay")

	response := h.urlService.CreateShortUrl(request.Url)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}
