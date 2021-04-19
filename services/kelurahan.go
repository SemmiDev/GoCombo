package services

import (
	"GoOrder/models"
	"context"
)

type KelurahanService struct {
	KelurahanRepo models.KelurahanRepository
}

// NewKelurahanService will initiate new usecase
func NewKelurahanService(kr models.KelurahanRepository) models.KelurahanService {
	return &KelurahanService{
		KelurahanRepo: kr,
	}
}

func (p *KelurahanService) GetByName(name string) (ID string, err error) {
	ID, err = p.KelurahanRepo.GetByName(name)
	return
}

func (p *KelurahanService) Joining(ctx context.Context) (res []models.Joining, err error) {
	res, err = p.KelurahanRepo.Joining(ctx)
	return
}

func (p *KelurahanService) GetByKodePos(kodepos string) (res *models.Kelurahan, err error) {
	res, err = p.KelurahanRepo.GetByKodePos(kodepos)
	return
}

func (p *KelurahanService) GetKelurahanByKecamatanID(ID string) (res []models.Kelurahan, err error) {
	res, err = p.KelurahanRepo.GetKelurahanByKecamatanID(ID)
	return res, err
}

func (p *KelurahanService) DeleteByID(ctx context.Context, ID string) error {
	err := p.KelurahanRepo.DeleteByID(ctx, ID)
	return err
}

func (p *KelurahanService) CreateKelurahan(ctx context.Context, kelurahan *models.Kelurahan) (id string, err error) {
	id, err = p.KelurahanRepo.Store(ctx, kelurahan)
	return
}

func (p *KelurahanService) UpdateKelurahan(ctx context.Context, ID string, kelurahan *models.Kelurahan) error {
	err := p.KelurahanRepo.UpdateById(
		ctx,
		ID,
		kelurahan,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *KelurahanService) GetAll() (res []models.Kelurahan, err error) {
	res, err = p.KelurahanRepo.GetAll()
	return
}

func (p *KelurahanService) GetByID(ID string) (res *models.Kelurahan, err error) {
	res, err = p.KelurahanRepo.GetByID(ID)
	return
}