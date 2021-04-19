package controllers

import (
	"GoOrder/models"
	httpUtils "GoOrder/utils/http"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
)

type kelurahanController struct {
	kelurahanService models.KelurahanService
	registrasiService models.RegistrasiService
}

// NewKelurahanController will initiate new api
func NewKelurahanController(r *mux.Router, ks models.KelurahanService, rs models.RegistrasiService) {
	kelurahanController := &kelurahanController{
		kelurahanService: ks,
		registrasiService: rs,
	}

	r.HandleFunc("/kelurahan", kelurahanController.GetAll).Methods("GET")
	r.HandleFunc("/kelurahan", kelurahanController.Create).Methods("POST")
	r.HandleFunc("/kelurahan/{id_kelurahan}", kelurahanController.GetByID).Methods("GET")
	r.HandleFunc("/kelurahan/kodepos/{kodepos}", kelurahanController.GetByKodePos).Methods("GET")
	r.HandleFunc("/kelurahan/kecamatan/{id_kecamatan}", kelurahanController.GetKelurahanByKecamatanID).Methods("GET")
	r.HandleFunc("/kelurahan/{id_kelurahan}", kelurahanController.UpdateByID).Methods("PUT")
	r.HandleFunc("/kelurahan/{id_kelurahan}", kelurahanController.DeleteByID).Methods("DELETE")
	r.HandleFunc("/join", kelurahanController.joining).Methods("GET")

	r.HandleFunc("/register", kelurahanController.ShowDataKelurahan)
	r.HandleFunc("/insertRegister", kelurahanController.InsertRegister)
	r.HandleFunc("/home", kelurahanController.Home)
}

var templateHtml = template.Must(template.ParseGlob("templates/kelurahan/*"))

func (c *kelurahanController) ShowDataKelurahan(w http.ResponseWriter, r *http.Request) {
	result, _ := c.kelurahanService.GetAll()
	_ = templateHtml.ExecuteTemplate(w, "Registration", result)
}

func (c kelurahanController) Home(w http.ResponseWriter, r *http.Request) {
	registers, _ := c.registrasiService.GetAll()
	_ = templateHtml.ExecuteTemplate(w, "View", registers)
}

func (c *kelurahanController) InsertRegister(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	result := r.Form["kelurahanchoice"]
	namaKelurahan := result[0]
	kelurahanByID, _ := c.kelurahanService.GetByName(namaKelurahan)

	log.Println(kelurahanByID)
	var register models.Registrasi

	register.RegistrasiID = uuid.NewV4().String()
	register.Nama = r.FormValue("fullname")
	register.Email = r.FormValue("email")
	register.Password = r.FormValue("password")
	register.Alamat = r.FormValue("address")
	register.KelurahanID = kelurahanByID

	_, _ = c.registrasiService.CreateRegistrasi(context.TODO(), &register)
	registers, _ := c.registrasiService.GetAll()
	_ = templateHtml.ExecuteTemplate(w, "View", registers)
}

func (c *kelurahanController) GetByKodePos(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kodepos := params["kodepos"]
	result, err := c.kelurahanService.GetByKodePos(kodepos)

	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get kelurahan by kode pos",
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

func (c *kelurahanController) GetAll(w http.ResponseWriter, r *http.Request) {
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

func (c *kelurahanController) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ID       		string `json:"id"`
		Kecamatan   	string `json:"id_kecamatan"`
		Nama       		string `json:"nama"`
		KodePos       	string `json:"kodepos"`

	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad r body", http.StatusBadRequest)
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

func (c *kelurahanController) GetByID(w http.ResponseWriter, r *http.Request) {
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

func (c *kelurahanController) GetKelurahanByKecamatanID(w http.ResponseWriter, r *http.Request) {
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

func (c *kelurahanController) UpdateByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kelurahanID := params["id_kelurahan"]

	var body struct {
		ID       		string `json:"id"`
		KecamatanID   	string `json:"id_kecamatan"`
		Nama       		string `json:"nama"`
		KodePos       	string `json:"kodepos"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad r body", http.StatusBadRequest)
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

func (c *kelurahanController) DeleteByID(w http.ResponseWriter, r *http.Request) {
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

func (c *kelurahanController) joining(w http.ResponseWriter, r *http.Request) {

	res, err := c.kelurahanService.Joining(context.TODO())
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to joining",
			http.StatusInternalServerError,
		)
		return
	}

	httpUtils.HandleJSONResponse(w,r,res)
}