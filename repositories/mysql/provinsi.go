package mysql

import (
	"GoOrder/models"
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"
)

type provinsiRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewProvinsiRepo will create provinsi rrepo
func NewProvinsiRepo(reader, writer *sql.DB) models.ProvinsiRepository {
	return &provinsiRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (p *provinsiRepo) GetByID(ID string) (res *models.Provinsi, err error) {
	query := sq.Select("*").
		From(PROVINSI).
		Where(sq.Eq{
			"id": ID,
	}).
	RunWith(p.Reader).
	PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var r models.Provinsi
		err = rows.Scan(
			&r.ProvinsiID,
			&r.Nama,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}
	return
}

func (p *provinsiRepo) DeleteByID(ctx context.Context, ID string) error {
	query := sq.Delete("").
		From(PROVINSI).
		Where(sq.Eq{
			"id": ID,
		}).
		RunWith(p.Reader)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *provinsiRepo) Store(ctx context.Context, provinsi *models.Provinsi) (provinsiID string, err error) {
	query := sq.Insert(PROVINSI).
		Columns(
			"id",
			"nama",
		).
		Values(
			provinsi.ProvinsiID,
			provinsi.Nama,
		).
		PlaceholderFormat(sq.Question)

	sqlInsert, argsInsert, err := query.ToSql()
	_, err = p.Writer.ExecContext(
		ctx,
		sqlInsert,
		argsInsert...,
	)

	if err != nil {
		return
	}

	provinsiID = provinsi.ProvinsiID
	return
}

func (p *provinsiRepo) UpdateById(ctx context.Context, ID string, provinsi *models.Provinsi) error {
	query := sq.Update(PROVINSI).
		Where(sq.Eq{
			"id": ID,
		}).
		SetMap(map[string]interface{}{
			"id":  provinsi.ProvinsiID,
			"nama": provinsi.Nama,
		}).
		RunWith(p.Writer).PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *provinsiRepo) GetAll() (res []models.Provinsi, err error) {
	query := sq.Select("*").
		From(PROVINSI).
		RunWith(p.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Provinsi
		err = rows.Scan(
			&r.ProvinsiID,
			&r.Nama,
		)

		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}
	return
}