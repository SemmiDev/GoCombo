package mysql

import (
	"GoOrder/models"
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"
	"log"
)

type kelurahanRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewKelurahanRepo will create kelurahan_kota rrepo
func NewKelurahanRepo(reader, writer *sql.DB) models.KelurahanRepository {
	return &kelurahanRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (k *kelurahanRepo) GetByName(name string) (ID string, err error) {
	query := sq.Select("id").
		From(KELURAHAN).
		Where(sq.Eq{
			"nama": name,
		}).
		RunWith(k.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var id string
		err = rows.Scan(
			&id,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		ID = id
	}
	return
}

func (k *kelurahanRepo) GetByKodePos(kodepos string) (res *models.Kelurahan, err error) {
	query := sq.Select("*").
		From(KELURAHAN).
		Where(sq.Eq{
			"kodepos": kodepos,
		}).
		RunWith(k.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var r models.Kelurahan
		err = rows.Scan(
			&r.KelurahanID,
			&r.KecamatanID,
			&r.Nama,
			&r.KodePos,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}
	return
}

func (k *kelurahanRepo) GetByID(ID string) (res *models.Kelurahan, err error) {
	query := sq.Select("*").
		From(KELURAHAN).
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
		var r models.Kelurahan
		err = rows.Scan(
			&r.KelurahanID,
			&r.KecamatanID,
			&r.Nama,
			&r.KodePos,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}
	return
}

func (k *kelurahanRepo) GetKelurahanByKecamatanID(ID string) (res []models.Kelurahan, err error) {

	query := sq.Select("*").
		From(KELURAHAN).
		Where(sq.Eq{
			"id_kecamatan": ID,
		}).
		RunWith(k.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Kelurahan
		err = rows.Scan(
			&r.KelurahanID,
			&r.KecamatanID,
			&r.Nama,
			&r.KodePos,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}
	return
}

func (k *kelurahanRepo) DeleteByID(ctx context.Context, ID string) error {
	query := sq.Delete("").
		From(KELURAHAN).
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

func (k *kelurahanRepo) Store(ctx context.Context, kelurahan *models.Kelurahan) (kelurahanID string, err error) {
	query := sq.Insert(KELURAHAN).
		Columns(
			"id",
			"id_kecamatan",
			"nama",
			"kodepos",
		).
		Values(
			kelurahan.KelurahanID,
			kelurahan.KecamatanID,
			kelurahan.Nama,
			kelurahan.KodePos,
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

	kelurahanID = kelurahan.KelurahanID
	return
}

func (k *kelurahanRepo) UpdateById(ctx context.Context, ID string, kelurahan *models.Kelurahan) error {
	query := sq.Update(KELURAHAN).
		Where(sq.Eq{
			"id": ID,
		}).
		SetMap(map[string]interface{}{
			"id": ID,
			"id_kecamatan": kelurahan.KecamatanID,
			"nama": kelurahan.Nama,
			"kodepos": kelurahan.KodePos,
		}).
		RunWith(k.Writer).PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (k *kelurahanRepo) GetAll() (res []models.Kelurahan, err error) {
	query := sq.Select("*").
		From(KELURAHAN).
		RunWith(k.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Kelurahan
		err = rows.Scan(
			&r.KelurahanID,
			&r.KecamatanID,
			&r.Nama,
			&r.KodePos,
		)

		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}
	return
}

func (r *kelurahanRepo) Joining(ctx context.Context) (res []models.Joining, err error) {

	rows, err := r.Reader.QueryContext(ctx, "" +
		"SELECT registrasi.id," +
		"\n registrasi.nama," +
		"\n registrasi.email," +
		"\n registrasi.alamat," +
		"\n kelurahan.nama," +
		"\n kelurahan.kodepos," +
		"\n kecamatan.nama," +
		"\n kabupaten_kota.nama," +
		"\n provinsi.nama" +
		"\n FROM registrasi JOIN kelurahan ON registrasi.id_kelurahan = kelurahan.id" +
		"\n JOIN kecamatan on kelurahan.id_kecamatan = kecamatan.id" +
		"\n JOIN kabupaten_kota on kecamatan.id_kabupaten_kota = kabupaten_kota.id" +
		"\n JOIN provinsi on kabupaten_kota.id_provinsi = provinsi.id")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var r models.Joining
		err = rows.Scan(
			&r.ID,
			&r.Name,
			&r.Email,
			&r.Address,
			&r.Kelurahan,
			&r.KodePos,
			&r.Kecamatan,
			&r.Kabupaten,
			&r.Provinsi,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}
	return
}