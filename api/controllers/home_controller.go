package controllers

import (
	"net/http"

	"github.com/victorsteven/fullstack/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, 0, "Success", "Qtracking - Golang Api")

}
