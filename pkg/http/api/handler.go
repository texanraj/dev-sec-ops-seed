package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

type requestHandler struct {
}

func NewAPIHandler() http.Handler {
	handler := &requestHandler{}
	router := mux.NewRouter()
	v1Router := router.PathPrefix("/api").Subrouter()

	v1Router.Methods(http.MethodGet).Path("/health").HandlerFunc(handler.GetInfo)
	return router
}

func (h *requestHandler) GetInfo(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
}
