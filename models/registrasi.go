package models

import "context"

// Registrasi for registrasi table
type Registrasi struct {
	RegistrasiID 	string `json:"id"`
	Email			string `json:"email"`
	Password		string `json:"password"`
	Nama       		string `json:"nama"`
	Alamat      	string `json:"alamat"`
	KelurahanID 	string `json:"id_kelurahan"`
}

type UpdateRegister struct {
	RegistrasiID 	string `json:"id"`
	Email			string `json:"email"`
	Password		string `json:"password"`
	Nama       		string `json:"nama"`
	Alamat      	string `json:"alamat"`
	NamaKelurahan   []string `json:"nama_kelurahan"`
}

// RegistrasiRepository for repo
type RegistrasiRepository interface {
	GetAll() ([]Registrasi, error)
	GetByID(ID string) (*Registrasi, error)
	GetRegistrasiyKelurahanID(ID string) ([]Registrasi,error)
	DeleteByID(ctx context.Context, ID string) error
	Store(ctx context.Context, registrasi *Registrasi) (string, error)
	UpdateById(ctx context.Context, ID string, registrasi *Registrasi) error
}

// RegistrasiService for service
type RegistrasiService interface {
	GetAll() ([]Registrasi, error)
	GetByID(ID string) (*Registrasi, error)
	GetRegistrasiyKelurahanID(ID string) ([]Registrasi,error)
	DeleteByID(ctx context.Context, ID string) error
	CreateRegistrasi(ctx context.Context, registrasi *Registrasi) (string, error)
	UpdateRegistrasi(ctx context.Context, ID string, registrasi *Registrasi) error
}