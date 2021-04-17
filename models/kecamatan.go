package models

import "context"

// Kecamatan for kecamatan table
type Kecamatan struct {
	KecamatanID 	string `json:"id"`
	KabupatenKotaID string `json:"id_kabupaten_kota"`
	Nama       		string `json:"nama"`
}

// KecamatanRepository for repo
type KecamatanRepository interface {
	GetAll() ([]Kecamatan, error)
	GetByID(ID string) (*Kecamatan, error)
	GetKecamatanByKabupatenID(ID string) ([]Kecamatan,error)
	DeleteByID(ctx context.Context, ID string) error
	Store(ctx context.Context, kecamatan *Kecamatan) (string, error)
	UpdateById(ctx context.Context, ID string, kecamatan *Kecamatan) error
}

// KecamatanService for service
type KecamatanService interface {
	GetAll() ([]Kecamatan, error)
	GetByID(ID string) (*Kecamatan, error)
	DeleteByID(ctx context.Context, ID string) error
	GetKecamatanByKabupatenID(ID string) ([]Kecamatan,error)
	CreateKecamatan(ctx context.Context, kecamatan *Kecamatan) (string, error)
	UpdateKecamatan(ctx context.Context, ID string, kecamatan *Kecamatan) error
}