package mysql

import (
	"GoOrder/models"
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"
)

type kecamatanRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewKabupatenKotaRepo will create kabupaten_kota rrepo
func NewKecamatanRepo(reader, writer *sql.DB) models.KecamatanRepository {
	return &kecamatanRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (k *kecamatanRepo) GetByID(ID string) (res *models.Kecamatan, err error) {
	query := sq.Select("*").
		From(KECAMATAN).
		Where(sq.Eq{
			"id": ID,
		}).
		RunWith(k.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var r models.Kecamatan
		err = rows.Scan(
			&r.KecamatanID,
			&r.Nama,
			&r.KabupatenKotaID,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}
	return
}

func (k *kecamatanRepo) GetKecamatanByKabupatenID(ID string) (res []models.Kecamatan, err error) {
	query := sq.Select("*").
		From(KECAMATAN).
		Where(sq.Eq{
			"id_kabupaten_kota": ID,
		}).
		RunWith(k.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Kecamatan
		err = rows.Scan(
			&r.KecamatanID,
			&r.Nama,
			&r.KabupatenKotaID,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}
	return
}

func (k *kecamatanRepo) DeleteByID(ctx context.Context, ID string) error {
	query := sq.Delete("").
		From(KECAMATAN).
		Where(sq.Eq{
			"id": ID,
		}).
		RunWith(k.Reader)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (k *kecamatanRepo) Store(ctx context.Context, kecamatan *models.Kecamatan) (kecamatanID string, err error) {
	query := sq.Insert(KECAMATAN).
		Columns(
			"id",
			"nama",
			"id_kabupaten_kota",
		).
		Values(
			kecamatan.KecamatanID,
			kecamatan.Nama,
			kecamatan.KabupatenKotaID,
		).
		PlaceholderFormat(sq.Question)

	sqlInsert, argsInsert, err := query.ToSql()
	_, err = k.Writer.ExecContext(
		ctx,
		sqlInsert,
		argsInsert...,
	)

	if err != nil {
		return
	}

	kecamatanID = kecamatan.KecamatanID
	return
}

func (k *kecamatanRepo) UpdateById(ctx context.Context, ID string, kecamatan *models.Kecamatan) error {
	query := sq.Update(KECAMATAN).
		Where(sq.Eq{
			"id": ID,
		}).
		SetMap(map[string]interface{}{
			"id": kecamatan.KecamatanID,
			"nama": kecamatan.Nama,
			"id_kabupaten_kota": kecamatan.KabupatenKotaID,
		}).
		RunWith(k.Writer).PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (k *kecamatanRepo) GetAll() (res []models.Kecamatan, err error) {
	query := sq.Select("*").
		From(KECAMATAN).
		RunWith(k.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Kecamatan
		err = rows.Scan(
			&r.KecamatanID,
			&r.KabupatenKotaID,
			&r.Nama,
		)

		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}
	return
}