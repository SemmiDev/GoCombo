package controllers

import (
	"github.com/SemmiDev/go-combo/api/responses"
	"net/http"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To Go Combo API")
}
