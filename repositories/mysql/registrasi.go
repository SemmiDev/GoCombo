package mysql

import (
	"GoOrder/models"
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"
)

type registrasiRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewKelurahanRepo will create kelurahan_kota rrepo
func NewRegistrasiRepo(reader, writer *sql.DB) models.RegistrasiRepository {
	return &registrasiRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (r *registrasiRepo) GetRegistrasiyKelurahanID(ID string) (res []models.Registrasi, err error) {
	query := sq.Select("*").
		From(REGISTRASI).
		Where(sq.Eq{
			"id_kelurahan": ID,
		}).
		RunWith(r.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Registrasi
		err = rows.Scan(
			&r.RegistrasiID,
			&r.Email,
			&r.Password,
			&r.Nama,
			&r.Alamat,
			&r.KelurahanID,
		)
		logger.Println(r)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}
	return
}

func (r *registrasiRepo) GetByID(ID string) (res *models.Registrasi, err error) {
	query := sq.Select("*").
		From(REGISTRASI).
		Where(sq.Eq{
			"id": ID,
		}).
		RunWith(r.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var r models.Registrasi
		err = rows.Scan(
			&r.RegistrasiID,
			&r.Email,
			&r.Password,
			&r.Nama,
			&r.Alamat,
			&r.KelurahanID,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}
	return
}

func (r *registrasiRepo) DeleteByID(ctx context.Context, ID string) error {
	query := sq.Delete("").
		From(REGISTRASI).
		Where(sq.Eq{
			"id": ID,
		}).
		RunWith(r.Reader)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *registrasiRepo) Store(ctx context.Context, registrasi *models.Registrasi) (registrasiID string, err error) {
	query := sq.Insert(REGISTRASI).
		Columns(
			"id",
			"email",
			"password",
			"nama",
			"alamat",
			"id_kelurahan",
		).
		Values(
			registrasi.RegistrasiID,
			registrasi.Email,
			registrasi.Password,
			registrasi.Nama,
			registrasi.Alamat,
			registrasi.KelurahanID,
		).
		PlaceholderFormat(sq.Question)

	sqlInsert, argsInsert, err := query.ToSql()
	_, err = r.Writer.ExecContext(
		ctx,
		sqlInsert,
		argsInsert...,
	)

	if err != nil {
		return
	}

	registrasiID = registrasi.RegistrasiID
	return
}

func (r *registrasiRepo) UpdateById(ctx context.Context, ID string, registrasi *models.Registrasi) error {
	query := sq.Update(REGISTRASI).
		Where(sq.Eq{
			"id": ID,
		}).
		SetMap(map[string]interface{}{
			"email": registrasi.Nama,
			"password": registrasi.Password,
			"nama": registrasi.Nama,
			"alamat": registrasi.Alamat,
			"id_kelurahan": registrasi.KelurahanID,
		}).
		RunWith(r.Writer).PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *registrasiRepo) GetAll() (res []models.Registrasi, err error) {
	query := sq.Select("*").
		From(REGISTRASI).
		RunWith(r.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Registrasi
		err = rows.Scan(
			&r.RegistrasiID,
			&r.Email,
			&r.Password,
			&r.Nama,
			&r.Alamat,
			&r.KelurahanID,
		)

		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}
	return
}