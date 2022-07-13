package repository

import (
	"context"
	"database/sql"
	"errors"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"programmerzamannow/belajar-golang-restful-api/model/domain"
)

type HotelRepositoryImpl struct {
}

func NewHotelRepository() HotelRepository {
	return &HotelRepositoryImpl{}
}

func (repository *HotelRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, hotel domain.Hotel) domain.Hotel {
	SQL := "insert into hotel(name, address, province, city, zip_code, star) values (?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, hotel.Name, hotel.Address, hotel.Province, hotel.City, hotel.ZipCode, hotel.Star)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	hotel.Id = int(id)
	return hotel
}

func (repository *HotelRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, hotel domain.Hotel) domain.Hotel {
	SQL := "update hotel set name = ?, address = ?, province = ?, city = ?, zip_code = ?, star = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, hotel.Name, hotel.Address, hotel.Province, hotel.City, hotel.ZipCode, hotel.Star, hotel.Id)
	helper.PanicIfError(err)

	return hotel
}

func (repository *HotelRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, hotel domain.Hotel) {
	SQL := "delete from hotel where id = ?"
	_, err := tx.ExecContext(ctx, SQL, hotel.Id)
	helper.PanicIfError(err)
}

func (repository *HotelRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, hotelId int) (domain.Hotel, error) {
	SQL := "select id, name, address, province, city, zip_code, star from hotel where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, hotelId)
	helper.PanicIfError(err)
	defer rows.Close()

	hotel := domain.Hotel{}
	if rows.Next() {
		err := rows.Scan(&hotel.Id, &hotel.Name, &hotel.Address, &hotel.Province, &hotel.City, &hotel.ZipCode, &hotel.Star)
		helper.PanicIfError(err)
		return hotel, nil
	} else {
		return hotel, errors.New("hotel is not found")
	}
}

func (repository *HotelRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Hotel {
	SQL := "select id, name, address, province, city, zip_code, star from hotel"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var hotels []domain.Hotel
	for rows.Next() {
		hotel := domain.Hotel{}
		err := rows.Scan(&hotel.Id, &hotel.Name, &hotel.Address, &hotel.Province, &hotel.City, &hotel.ZipCode, &hotel.Star)
		helper.PanicIfError(err)
		hotels = append(hotels, hotel)
	}
	return hotels
}
