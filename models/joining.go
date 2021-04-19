package models

type Joining struct {
	Email 		string `json:"email"`
	Kelurahan	string `json:"kelurahan"`
	KodePos    	string `json:"kode_pos"`
	Kecamatan   string `json:"kecamatan"`
	Kabupaten   string `json:"kabupaten"`
	Provinsi    string `json:"provinsi"`
}