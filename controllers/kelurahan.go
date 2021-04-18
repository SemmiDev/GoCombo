package controllers

import (
	"GoOrder/models"
	httpUtils "GoOrder/utils/http"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type kelurahanController struct {
	kelurahanService models.KelurahanService
}

// NewKelurahanController will initiate new api
func NewKelurahanController(r *mux.Router, cuc models.KelurahanService) {
	kelurahanController := &kelurahanController{
		kelurahanService: cuc,
	}

	r.HandleFunc("/kelurahan", kelurahanController.GetAll).Methods("GET")
	r.HandleFunc("/kelurahan", kelurahanController.Create).Methods("POST")
	r.HandleFunc("/kelurahan/{id_kelurahan}", kelurahanController.GetByID).Methods("GET")
	r.HandleFunc("/kelurahan/kecamatan/{id_kecamatan}", kelurahanController.GetKelurahanByKecamatanID).Methods("GET")
	r.HandleFunc("/kelurahan/{id_kelurahan}", kelurahanController.UpdateByID).Methods("PUT")
	r.HandleFunc("/kelurahan/{id_kelurahan}", kelurahanController.DeleteByID).Methods("DELETE")
}

func (c kelurahanController) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := c.kelurahanService.GetAll()
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get kelurahan data",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data []models.Kelurahan `json:"data"`
	}
	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (c kelurahanController) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ID       		string `json:"id"`
		Kecamatan   	string `json:"id_kecamatan"`
		Nama       		string `json:"nama"`
		KodePos       	string `json:"kodepos"`

	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}

	defer r.Body.Close()

	kelurahan := models.Kelurahan{
		KelurahanID: 	body.ID,
		KecamatanID:    body.Kecamatan,
		Nama:        	body.Nama,
		KodePos:        body.KodePos,
	}

	id, err := c.kelurahanService.CreateKelurahan(context.TODO(), &kelurahan)
	if err != nil {
		httpUtils.HandleError(w, r, err, "failed to create kelurahan", http.StatusInternalServerError)
		return
	}

	var response struct {
		KelurahanID string `json:"id_kelurahan"`
	}

	response.KelurahanID = id
	httpUtils.HandleJSONResponse(w, r, response)
}

func (c kelurahanController) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kelurahanID := params["id_kelurahan"]
	result, err := c.kelurahanService.GetByID(kelurahanID)

	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get kelurahan by id",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data *models.Kelurahan `json:"data"`
	}

	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (c kelurahanController) GetKelurahanByKecamatanID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kecamatanID := params["id_kecamatan"]

	result, err := c.kelurahanService.GetKelurahanByKecamatanID(kecamatanID)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get kelurahan by kecamatan ID",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data []models.Kelurahan `json:"data"`
	}

	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (c kelurahanController) UpdateByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kelurahanID := params["id_kelurahan"]

	var body struct {
		ID       		string `json:"id"`
		KecamatanID   	string `json:"id_kecamatan"`
		Nama       		string `json:"nama"`
		KodePos       	string `json:"kodepos"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	kelurahan := models.Kelurahan{
		KelurahanID: body.ID,
		KecamatanID: body.KecamatanID,
		Nama: body.Nama,
		KodePos: body.KodePos,
	}

	err := c.kelurahanService.UpdateKelurahan(context.TODO(), kelurahanID, &kelurahan)

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

func (c kelurahanController) DeleteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kelurahanID := params["id_kelurahan"]

	errdel := c.kelurahanService.DeleteByID(context.TODO(), kelurahanID)
	if errdel != nil {
		httpUtils.HandleError(
			w,
			r,
			errdel,
			"failed to delete kelurahan by id",
			http.StatusInternalServerError,
		)
		return
	}

	httpUtils.HandleNoJSONResponse(w)
}