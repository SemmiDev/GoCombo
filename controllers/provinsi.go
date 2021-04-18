package controllers

import (
	"GoOrder/models"
	httpUtils "GoOrder/utils/http"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type provinsiController struct {
	provinsiService models.ProvinsiService
}

// NewKabupatenController will initiate new api
func NewProvinsiController(r *mux.Router, cuc models.ProvinsiService) {
	provinsiController := &provinsiController{
		provinsiService: cuc,
	}

	r.HandleFunc("/provinsi", provinsiController.GetAll).Methods("GET")
	r.HandleFunc("/provinsi", provinsiController.Create).Methods("POST")
	r.HandleFunc("/provinsi/{id_provinsi}", provinsiController.GetByID).Methods("GET")
	r.HandleFunc("/provinsi/{id_provinsi}", provinsiController.UpdateByID).Methods("PUT")
	r.HandleFunc("/provinsi/{id_provinsi}", provinsiController.DeleteByID).Methods("DELETE")
}


func (c provinsiController) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := c.provinsiService.GetAll()
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get provinsi data",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data []models.Provinsi `json:"data"`
	}
	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (c provinsiController) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ID       		string `json:"id"`
		Nama       		string `json:"nama"`

	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}

	defer r.Body.Close()

	provinsi := models.Provinsi{
		ProvinsiID: body.ID,
		Nama:		body.Nama,
	}

	id, err := c.provinsiService.CreateProvinsi(context.TODO(), &provinsi)
	if err != nil {
		httpUtils.HandleError(w, r, err, "failed to create provinsi", http.StatusInternalServerError)
		return
	}

	var response struct {
		ProvinsiID string `json:"id_provinsi"`
	}

	response.ProvinsiID = id
	httpUtils.HandleJSONResponse(w, r, response)
}

func (c provinsiController) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	provinsiID := params["id_provinsi"]
	result, err := c.provinsiService.GetByID(provinsiID)

	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get provinsi by id",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data *models.Provinsi `json:"data"`
	}

	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (c provinsiController) UpdateByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	provinsiID := params["id_provinsi"]

	var body struct {
		ID       		string `json:"id"`
		Nama       		string `json:"nama"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	provinsi := models.Provinsi{
		ProvinsiID: body.ID,
		Nama: body.Nama,
	}

	err := c.provinsiService.UpdateProvinsi(context.TODO(), provinsiID, &provinsi)

	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			params["status"]+" failed",
			http.StatusBadRequest,
		)
		return
	}

	httpUtils.HandleJSONResponse(w, r, err)
}

func (c provinsiController) DeleteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	provinsiID := params["id_provinsi"]

	errdel := c.provinsiService.DeleteByID(context.TODO(), provinsiID)
	if errdel != nil {
		httpUtils.HandleError(
			w,
			r,
			errdel,
			"failed to delete provinsi by id",
			http.StatusInternalServerError,
		)
		return
	}

	httpUtils.HandleNoJSONResponse(w)
}
