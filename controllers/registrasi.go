package controllers

import (
	"GoOrder/models"
	httpUtils "GoOrder/utils/http"
	"GoOrder/utils/passwords"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type registrasiController struct {
	registrasiService models.RegistrasiService
}

// NewRegistrasiController will initiate new api
func NewRegistrasiController(r *mux.Router, cuc models.RegistrasiService) {
	registrasiController := &registrasiController{
		registrasiService: cuc,
	}

	r.HandleFunc("/registrasi", registrasiController.GetAll).Methods("GET")
	r.HandleFunc("/registrasi", registrasiController.Create).Methods("POST")
	r.HandleFunc("/registrasi/{id_registrasi}", registrasiController.GetByID).Methods("GET")
	r.HandleFunc("/registrasi/kelurahan/{id_kelurahan}", registrasiController.GetByKelurahanID).Methods("GET")
	r.HandleFunc("/registrasi/{id_registrasi}", registrasiController.UpdateByID).Methods("PUT")
	r.HandleFunc("/registrasi/{id_registrasi}", registrasiController.DeleteByID).Methods("DELETE")
}


func (c registrasiController) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := c.registrasiService.GetAll()
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get registrasi data",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data []models.Registrasi `json:"data"`
	}
	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (c registrasiController) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email       	string `json:"email"`
		Password    	string `json:"password"`
		Nama    		string `json:"nama"`
		Alamat    		string `json:"alamat"`
		KelurahanID    	string `json:"id_kelurahan"`

	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	register := models.Registrasi{
		RegistrasiID: uuid.NewV4().String(),
		Email:        body.Email,
		Password:     body.Password,
		Nama:         body.Nama,
		Alamat:       body.Alamat,
		KelurahanID:  body.KelurahanID,
	}

	id, err := c.registrasiService.CreateRegistrasi(context.TODO(), &register)
	if err != nil {
		httpUtils.HandleError(w, r, err, "failed to create register", http.StatusInternalServerError)
		return
	}

	var response struct {
		registerID string `json:"id_registrasi"`
	}

	response.registerID = id

	httpUtils.HandleJSONResponse(w, r, response)
}

func (c registrasiController) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	registrasiID := params["id_registrasi"]
	result, err := c.registrasiService.GetByID(registrasiID)

	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get registrasi by id",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data *models.Registrasi `json:"data"`
	}

	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (c registrasiController) GetByKelurahanID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kelurahanID := params["id_kelurahan"]

	result, err := c.registrasiService.GetRegistrasiyKelurahanID(kelurahanID)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get registers by kelurahan ID",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data []models.Registrasi `json:"data"`
	}

	data.Data = result

	httpUtils.HandleJSONResponse(w, r, data)
}

func (c registrasiController) UpdateByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	registrasiID := params["id_registrasi"]

	var body struct {
		Email       	string `json:"email"`
		Password    	string `json:"password"`
		Nama    		string `json:"nama"`
		Alamat    		string `json:"alamat"`
		KelurahanID    	string `json:"id_kelurahan"`

	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	cust := models.Registrasi{
		Email:        body.Email,
		Password:     body.Password,
		Nama:         body.Nama,
		Alamat:       body.Alamat,
		KelurahanID:  body.KelurahanID,
	}

	// Hashing first
	passwords.HashAndSaltPassword(&cust.Password)
	err := c.registrasiService.UpdateRegistrasi(context.TODO(), registrasiID, &cust)

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

func (c registrasiController) DeleteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	registrasiID := params["id_registrasi"]

	errdel := c.registrasiService.DeleteByID(context.TODO(), registrasiID)
	if errdel != nil {
		httpUtils.HandleError(
			w,
			r,
			errdel,
			"failed to delete registrasi by id",
			http.StatusInternalServerError,
		)
		return
	}

	httpUtils.HandleNoJSONResponse(w)
}