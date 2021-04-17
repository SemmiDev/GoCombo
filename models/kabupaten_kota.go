package models

import "context"

// KabupatenKota for provinsi table
type KabupatenKota struct {
	KabupatenKotaID string `json:"id"`
	ProvinsiID    	string `json:"id_provinsi"`
	Nama          	string `json:"nama"`
}

// KabupatenKotaRepository for repo
type KabupatenKotaRepository interface {
	GetAll() ([]KabupatenKota, error)
	GetByID(ID string) (*KabupatenKota, error)
	UpdateById(ctx context.Context, ID string, kabupatenKota *KabupatenKota) error
	GetByProvinsiID(ID string) ([]KabupatenKota, error)
	DeleteByID(ctx context.Context, ID string) error
	Store(ctx context.Context, kabupatenKota *KabupatenKota) (string, error)
}

// KabupatenKotaService for service
type KabupatenKotaService interface {
	GetAll() ([]KabupatenKota, error)
	GetByID(ID string) (*KabupatenKota, error)
	UpdateById(ctx context.Context, ID string, kabupatenKota *KabupatenKota) error
	DeleteByID(ctx context.Context, ID string) error
	GetKabupatenByProvinsiID(ID string) ([]KabupatenKota, error)
	CreateKabupatenKota(ctx context.Context, kabupatenKota *KabupatenKota) (string, error)
}