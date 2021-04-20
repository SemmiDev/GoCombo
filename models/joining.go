package models

type Joining struct {
	ID 			string `json:"id"`
	Name 		string `json:"name"`
	Email 		string `json:"email"`
	Address 	string `json:"address"`
	Kelurahan	string `json:"kelurahan"`
	KodePos    	string `json:"kode_pos"`
	Kecamatan   string `json:"kecamatan"`
	Kabupaten   string `json:"kabupaten"`
	Provinsi    string `json:"provinsi"`
}