package services

import (
	"GoOrder/models"
	"context"
)

type ProvinsiService struct {
	ProvinsiRepo models.ProvinsiRepository
}

// NewProvinsiService will initiate new usecase
func NewProvinsiService(rr models.ProvinsiRepository) models.ProvinsiService {
	return &ProvinsiService{
		ProvinsiRepo: rr,
	}
}

func (p *ProvinsiService) GetByID(ID string) (res *models.Provinsi, err error) {
	res, err = p.ProvinsiRepo.GetByID(ID)
	return
}

func (p *ProvinsiService) DeleteByID(ctx context.Context, ID string) error {
	err := p.ProvinsiRepo.DeleteByID(ctx, ID)
	return err
}

func (p *ProvinsiService) CreateProvinsi(ctx context.Context, provinsi *models.Provinsi) (id string, err error) {
	id, err = p.ProvinsiRepo.Store(ctx, provinsi)
	return
}

func (p *ProvinsiService) UpdateProvinsi(ctx context.Context, ID string, provinsi *models.Provinsi) error {
	err := p.ProvinsiRepo.UpdateById(
		ctx,
		ID,
		provinsi,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProvinsiService) GetAll() (res []models.Provinsi, err error) {
	res, err = p.ProvinsiRepo.GetAll()
	return
}