package repository

import (
	"context"
	"database/sql"
	"programmerzamannow/belajar-golang-restful-api/model/domain"
	"programmerzamannow/belajar-golang-restful-api/model/web"
)

type HotelRepository interface {
	Save(ctx context.Context, tx *sql.Tx, hotel domain.Hotel) domain.Hotel
	Update(ctx context.Context, tx *sql.Tx, hotel domain.Hotel) domain.Hotel
	Delete(ctx context.Context, tx *sql.Tx, hotel domain.Hotel)
	FindById(ctx context.Context, tx *sql.Tx, hotelId int) (web.HotelResponse, error)
	FindAll(ctx context.Context, tx *sql.Tx) []web.HotelResponse
}
