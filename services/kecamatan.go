package services

import (
	"GoOrder/models"
	"context"
)

type kecamatanService struct {
	KecamatanRepo models.KecamatanRepository
}

// NewKabupatenKotaService will initiate new usecase
func NewKecamatanService(rr models.KecamatanRepository) models.KecamatanService {
	return &kecamatanService{
		KecamatanRepo: rr,
	}
}

func (r *kecamatanService) DeleteByID(ctx context.Context, ID string) error {
	err := r.KecamatanRepo.DeleteByID(ctx, ID)
	return err
}

func (r *kecamatanService) GetAll() (res []models.Kecamatan, err error) {
	res, err = r.KecamatanRepo.GetAll()
	return
}

func (r *kecamatanService) GetByID(ID string) (res *models.Kecamatan, err error) {
	res, err = r.KecamatanRepo.GetByID(ID)
	return
}

func (r *kecamatanService) GetKecamatanByKabupatenID(ID string) (res []models.Kecamatan, err error) {
	res, err = r.KecamatanRepo.GetKecamatanByKabupatenID(ID)
	return res, err
}

func (r *kecamatanService) CreateKecamatan(ctx context.Context, kecamatan *models.Kecamatan) (id string, err error) {
	id, err = r.KecamatanRepo.Store(ctx, kecamatan)
	return
}

func (r *kecamatanService) UpdateKecamatan(ctx context.Context, ID string, kecamatan *models.Kecamatan) error {
	err := r.KecamatanRepo.UpdateById(
		ctx,
		ID,
		kecamatan,
	)

	if err != nil {
		return err
	}

	return nil
}