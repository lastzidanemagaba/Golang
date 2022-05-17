package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/victorsteven/fullstack/api/models"
	"github.com/victorsteven/fullstack/api/responses"
)

func (server *Server) GetNegaras(w http.ResponseWriter, r *http.Request) {

	user := models.Cr_Negara{}

	users, err := user.FindAllNegaras(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, 0, "Success", users)
}

func (server *Server) GetNegara(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.Cr_Negara{}
	userGotten, err := user.FindNegaraByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, 0, "Success", userGotten)
}
