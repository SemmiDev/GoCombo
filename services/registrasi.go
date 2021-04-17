package services

import (
	"GoOrder/models"
	"GoOrder/utils/passwords"
	"context"
)

type registrasiService struct {
	RegistrasiRepo models.RegistrasiRepository
}

// NewRegistrasiService will initiate new usecase
func NewRegistrasiService(rr models.RegistrasiRepository) models.RegistrasiService {
	return &registrasiService{
		RegistrasiRepo: rr,
	}
}

func (r *registrasiService) GetAll() (res []models.Registrasi, err error) {
	res, err = r.RegistrasiRepo.GetAll()
	return
}

func (r *registrasiService) GetByID(ID string) (res *models.Registrasi, err error) {
	res, err = r.RegistrasiRepo.GetByID(ID)
	return
}

func (r *registrasiService) GetRegistrasiyKelurahanID(ID string) (res []models.Registrasi, err error) {
	res, err = r.RegistrasiRepo.GetRegistrasiyKelurahanID(ID)
	return res, err
}

func (r *registrasiService) DeleteByID(ctx context.Context, ID string) error {
	err := r.RegistrasiRepo.DeleteByID(ctx, ID)
	return err
}

func (r *registrasiService) CreateRegistrasi(ctx context.Context, registrasi *models.Registrasi) (id string, err error) {
	passwords.HashAndSaltPassword(&registrasi.Password)
	id, err = r.RegistrasiRepo.Store(ctx, registrasi)
	return
}

func (r *registrasiService) UpdateRegistrasi(ctx context.Context, ID string, registrasi *models.Registrasi) error {
	err := r.RegistrasiRepo.UpdateById(
		ctx,
		ID,
		registrasi,
	)

	if err != nil {
		return err
	}

	return nil
}