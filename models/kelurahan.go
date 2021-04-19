package models

import "context"

// Kelurahan for kelurahan table
type Kelurahan struct {
	KelurahanID string `json:"id"`
	KecamatanID string `json:"id_kecamatan"`
	Nama       	string `json:"nama"`
	KodePos    	string `json:"kodepos"`
}

// KelurahanRepository for repo
type KelurahanRepository interface {
	GetAll() ([]Kelurahan, error)
	GetByID(ID string) (*Kelurahan, error)
	GetByKodePos(kodepos string) (*Kelurahan, error)
	GetByName(name string) (string, error)
	GetKelurahanByKecamatanID(ID string) ([]Kelurahan,error)
	DeleteByID(ctx context.Context, ID string) error
	Store(ctx context.Context, kelurahan *Kelurahan) (string, error)
	UpdateById(ctx context.Context, ID string, kelurahan *Kelurahan) error
	Joining(ctx context.Context) ([]Joining, error)
}

// KelurahanService for service
type KelurahanService interface {
	GetAll() ([]Kelurahan, error)
	GetByID(ID string) (*Kelurahan, error)
	GetByKodePos(kodepos string) (*Kelurahan, error)
	GetByName(name string) (string, error)
	GetKelurahanByKecamatanID(ID string) ([]Kelurahan,error)
	DeleteByID(ctx context.Context, ID string) error
	CreateKelurahan(ctx context.Context, kelurahan *Kelurahan) (string, error)
	UpdateKelurahan(ctx context.Context, ID string, kelurahan *Kelurahan) error
	Joining(ctx context.Context) ([]Joining, error)
}