package mysql

import (
	"GoOrder/models"
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"
)

type kabupatenKotaRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewKabupatenKotaRepo will create kabupaten_kota rrepo
func NewKabupatenKotaRepo(reader, writer *sql.DB) models.KabupatenKotaRepository {
	return &kabupatenKotaRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (k *kabupatenKotaRepo) GetByID(ID string) (res *models.KabupatenKota, err error) {
	query := sq.Select("*").
		From(KABUPATEN_KOTA).
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
		var r models.KabupatenKota
		err = rows.Scan(
			&r.KabupatenKotaID,
			&r.Nama,
			&r.ProvinsiID,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}
	return
}

func (k *kabupatenKotaRepo) UpdateById(ctx context.Context, ID string, kabupatenKota *models.KabupatenKota) error {
	query := sq.Update(KABUPATEN_KOTA).
		Where(sq.Eq{
			"id": ID,
		}).
		SetMap(map[string]interface{}{
			"nama": kabupatenKota.Nama,
			"id_provinsi": kabupatenKota.ProvinsiID,
		}).
		RunWith(k.Writer).PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (k *kabupatenKotaRepo) GetByProvinsiID(ID string) (res []models.KabupatenKota, err error) {
	query := sq.Select("*").
		From(KABUPATEN_KOTA).
		Where(sq.Eq{
			"id_provinsi": ID,
		}).
		RunWith(k.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r models.KabupatenKota
		err = rows.Scan(
			&r.KabupatenKotaID,
			&r.Nama,
			&r.ProvinsiID,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}
	return
}

func (k *kabupatenKotaRepo) DeleteByID(ctx context.Context, ID string) error {
	query := sq.Delete("").
		From(KABUPATEN_KOTA).
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

func (k *kabupatenKotaRepo) Store(ctx context.Context, kabupatenKota *models.KabupatenKota) (kabupatenKotaID string, err error) {
	query := sq.Insert(KABUPATEN_KOTA).
		Columns(
			"id",
			"nama",
			"id_provinsi",
		).
		Values(
			kabupatenKota.KabupatenKotaID,
			kabupatenKota.Nama,
			kabupatenKota.ProvinsiID,
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

	kabupatenKotaID = kabupatenKota.KabupatenKotaID
	return
}

func (k *kabupatenKotaRepo) GetAll() (res []models.KabupatenKota, err error) {
	query := sq.Select("*").
		From(KABUPATEN_KOTA).
		RunWith(k.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r models.KabupatenKota
		err = rows.Scan(
			&r.KabupatenKotaID,
			&r.Nama,
			&r.ProvinsiID,
		)

		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}
	return
}