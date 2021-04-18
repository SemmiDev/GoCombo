package services

import (
	"GoOrder/models"
	"context"
)

type kabupatenKotaService struct {
	KabupatenKotaRepo models.KabupatenKotaRepository
}

// NewKabupatenKotaService will initiate new usecase
func NewKabupatenKotaService(rr models.KabupatenKotaRepository) models.KabupatenKotaService {
	return &kabupatenKotaService{
		KabupatenKotaRepo: rr,
	}
}

func (r *kabupatenKotaService) GetAll() (res []models.KabupatenKota, err error) {
	res, err = r.KabupatenKotaRepo.GetAll()
	return
}

func (r *kabupatenKotaService) GetByID(ID string) (res *models.KabupatenKota, err error) {
	res, err = r.KabupatenKotaRepo.GetByID(ID)
	return
}

func (r *kabupatenKotaService) UpdateById(ctx context.Context, ID string, kabupatenKota *models.KabupatenKota) error {
	err := r.KabupatenKotaRepo.UpdateById(
		ctx,
		ID,
		kabupatenKota,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *kabupatenKotaService) GetKabupatenByProvinsiID(ID string) (res []models.KabupatenKota, err error) {
	res, err = r.KabupatenKotaRepo.GetByProvinsiID(ID)
	return res, err
}

func (r *kabupatenKotaService) DeleteByID(ctx context.Context, ID string) error {
	err := r.KabupatenKotaRepo.DeleteByID(ctx, ID)
	return err
}

func (r *kabupatenKotaService) CreateKabupatenKota(ctx context.Context, kabupatenKota *models.KabupatenKota) (id string, err error) {
	id, err = r.KabupatenKotaRepo.Store(ctx, kabupatenKota)
	return
}

func (r *kabupatenKotaService) UpdateKabupatenKota(ctx context.Context, ID string, kabupatenKota *models.KabupatenKota) error {
	err := r.KabupatenKotaRepo.UpdateById(
		ctx,
		ID,
		kabupatenKota,
	)

	if err != nil {
		return err
	}

	return nil
}
