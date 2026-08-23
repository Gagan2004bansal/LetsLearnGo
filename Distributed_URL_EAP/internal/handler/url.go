package handler

import (
	"encoding/json"
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
	router.HandleFunc("/url/{shortcode}", h.GetUrl).Methods(http.MethodGet)
	router.HandleFunc("/url/{id}", h.DeleteUrl).Methods(http.MethodDelete)
}

func (h *UrlHandler) CreateUrl(w http.ResponseWriter, r *http.Request) {
	var request model.ReqUrl

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if request.Url == "" {
		http.Error(w, "please add url", http.StatusBadRequest)
		return
	}

	response := h.urlService.CreateShortUrl(request.Url)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}

func (h *UrlHandler) GetUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortcode := vars["shortcode"]

	if shortcode == "" {
		http.Error(w, "empty shorturl", http.StatusBadRequest)
		return
	}

	url, exists := h.urlService.GetLongUrl(shortcode)
	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url.Url, http.StatusFound)
}

func (h *UrlHandler) DeleteUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortcode := vars["id"]

	if shortcode == "" {
		http.Error(w, "empty shorturl", http.StatusBadRequest)
		return
	}

	res := h.urlService.DeleteShortUrl(shortcode)

	if !res {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}
