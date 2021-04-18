package controllers

import (
	"GoOrder/models"
	httpUtils "GoOrder/utils/http"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type kecamatanController struct {
	kecamatanService models.KecamatanService
}

// NewKecamatanController will initiate new api
func NewKecamatanController(r *mux.Router, cuc models.KecamatanService) {
	kecamatanController := &kecamatanController{
		kecamatanService: cuc,
	}

	r.HandleFunc("/kecamatan", kecamatanController.GetAll).Methods("GET")
	r.HandleFunc("/kecamatan", kecamatanController.Create).Methods("POST")
	r.HandleFunc("/kecamatan/{id_kecamatan}", kecamatanController.GetByID).Methods("GET")
	r.HandleFunc("/kecamatan/kabupaten/{id_kabupaten}", kecamatanController.GetByKabupatenID).Methods("GET")
	r.HandleFunc("/kecamatan/{id_kecamatan}", kecamatanController.UpdateByID).Methods("PUT")
	r.HandleFunc("/kecamatan/{id_kecamatan}", kecamatanController.DeleteByID).Methods("DELETE")
}

func (c kecamatanController) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := c.kecamatanService.GetAll()
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get kecamatan data",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data []models.Kecamatan `json:"data"`
	}
	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (c kecamatanController) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ID       		string `json:"id"`
		Nama       		string `json:"nama"`
		KabupatenKota   string `json:"id_kabupaten"`

	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}

	defer r.Body.Close()

	kecamatan := models.Kecamatan{
		KecamatanID: 		body.ID,
		Nama:        		body.Nama,
		KabupatenKotaID:    body.KabupatenKota,
	}

	id, err := c.kecamatanService.CreateKecamatan(context.TODO(), &kecamatan)
	if err != nil {
		httpUtils.HandleError(w, r, err, "failed to create kecamatan", http.StatusInternalServerError)
		return
	}

	var response struct {
		KecamatanID string `json:"id_kecamatan"`
	}

	response.KecamatanID = id
	httpUtils.HandleJSONResponse(w, r, response)
}

func (c kecamatanController) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kecamatanID := params["id_kecamatan"]
	result, err := c.kecamatanService.GetByID(kecamatanID)

	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get kecamatan by id",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data *models.Kecamatan `json:"data"`
	}

	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (c kecamatanController) GetByKabupatenID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kabupatenID := params["id_kabupaten"]

	result, err := c.kecamatanService.GetKecamatanByKabupatenID(kabupatenID)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get kecamatans by kabupaten ID",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data []models.Kecamatan `json:"data"`
	}

	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (c kecamatanController) UpdateByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kecamatanID := params["id_kecamatan"]

	var body struct {
		ID       		string `json:"id"`
		Nama       		string `json:"nama"`
		KabupatenID    	string `json:"id_kabupaten"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	kecamatan := models.Kecamatan{
		KecamatanID: body.ID,
		Nama: body.Nama,
		KabupatenKotaID: body.KabupatenID,
	}

	err := c.kecamatanService.UpdateKecamatan(context.TODO(), kecamatanID, &kecamatan)

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

func (c kecamatanController) DeleteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kecamatanID := params["id_kecamatan"]

	errdel := c.kecamatanService.DeleteByID(context.TODO(), kecamatanID)
	if errdel != nil {
		httpUtils.HandleError(
			w,
			r,
			errdel,
			"failed to delete kecamatan by id",
			http.StatusInternalServerError,
		)
		return
	}

	httpUtils.HandleNoJSONResponse(w)
}