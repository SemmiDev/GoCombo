package controllers

import (
	"GoOrder/models"
	httpUtils "GoOrder/utils/http"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type kabupatenKotaController struct {
	kabupatenKotaService models.KabupatenKotaService
}

// NewKabupatenKotaController will initiate new api
func NewKabupatenKotaController(r *mux.Router, cuc models.KabupatenKotaService) {
	kabupatenKotaController := &kabupatenKotaController{
		kabupatenKotaService: cuc,
	}

	r.HandleFunc("/kabupaten", kabupatenKotaController.GetAll).Methods("GET")
	r.HandleFunc("/kabupaten", kabupatenKotaController.Create).Methods("POST")
	r.HandleFunc("/kabupaten/{id_kabupaten}", kabupatenKotaController.GetByID).Methods("GET")
	r.HandleFunc("/kabupaten/provinsi/{id_provinsi}", kabupatenKotaController.GetByProvinsiID).Methods("GET")
	r.HandleFunc("/kabupaten/{id_kabupaten}", kabupatenKotaController.UpdateByID).Methods("PUT")
	r.HandleFunc("/kabupaten/{id_kabupaten}", kabupatenKotaController.DeleteByID).Methods("DELETE")
}


func (c *kabupatenKotaController) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := c.kabupatenKotaService.GetAll()
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get kabupaten data",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data []models.KabupatenKota `json:"data"`
	}
	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (c *kabupatenKotaController) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ID       		string `json:"id"`
		Nama       		string `json:"nama"`
		ProvinsiID    	string `json:"id_provinsi"`

	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}

	defer r.Body.Close()

	kabupatenKota := models.KabupatenKota{
		KabupatenKotaID: body.ID,
		Nama:        	body.Nama,
		ProvinsiID:     body.ProvinsiID,
	}

	id, err := c.kabupatenKotaService.CreateKabupatenKota(context.TODO(), &kabupatenKota)
	if err != nil {
		httpUtils.HandleError(w, r, err, "failed to create kabupatenKota", http.StatusInternalServerError)
		return
	}

	var response struct {
		KabupatenKotaID string `json:"id_kabupaten"`
	}

	response.KabupatenKotaID = id
	httpUtils.HandleJSONResponse(w, r, response)
}

func (c *kabupatenKotaController) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kabupatenID := params["id_kabupaten"]
	result, err := c.kabupatenKotaService.GetByID(kabupatenID)

	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get kabupaten by id",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data *models.KabupatenKota `json:"data"`
	}

	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (c *kabupatenKotaController) GetByProvinsiID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	provinsiID := params["id_provinsi"]

	result, err := c.kabupatenKotaService.GetKabupatenByProvinsiID(provinsiID)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get kabupatenKotas by provinsi ID",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data []models.KabupatenKota `json:"data"`
	}

	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (c *kabupatenKotaController) UpdateByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kabupatenID := params["id_kabupaten"]

	var body struct {
		ID       		string `json:"id"`
		Nama       		string `json:"nama"`
		ProvinsiID    	string `json:"id_provinsi"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	kabupaten := models.KabupatenKota{
		Nama: body.Nama,
		ProvinsiID: body.ProvinsiID,
	}

	err := c.kabupatenKotaService.UpdateById(context.TODO(), kabupatenID, &kabupaten)

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

func (c *kabupatenKotaController) DeleteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kabupatenID := params["id_kabupaten"]

	errdel := c.kabupatenKotaService.DeleteByID(context.TODO(), kabupatenID)
	if errdel != nil {
		httpUtils.HandleError(
			w,
			r,
			errdel,
			"failed to delete kabupaten by id",
			http.StatusInternalServerError,
		)
		return
	}

	httpUtils.HandleNoJSONResponse(w)
}