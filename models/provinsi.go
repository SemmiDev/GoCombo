package models

import "context"

// Provinsi for provinsi table
type Provinsi struct {
	ProvinsiID string `json:"id"`
	Nama       string `json:"nama"`
}

// ProvinsiRepository for repo
type ProvinsiRepository interface {
	GetAll() ([]Provinsi, error)
	GetByID(ID string) (*Provinsi, error)
	DeleteByID(ctx context.Context, ID string) error
	Store(ctx context.Context, provinsi *Provinsi) (string, error)
	UpdateById(ctx context.Context, ID string, provinsi *Provinsi) error
}
// ProvinsiService for service
type ProvinsiService interface {
	GetAll() ([]Provinsi, error)
	GetByID(ID string) (*Provinsi, error)
	DeleteByID(ctx context.Context, ID string) error
	CreateProvinsi(ctx context.Context, provinsi *Provinsi) (string, error)
	UpdateProvinsi(ctx context.Context, ID string, provinsi *Provinsi) error
}