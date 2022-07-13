package service

import (
	"context"
	"programmerzamannow/belajar-golang-restful-api/model/web"
)

type UserHotelService interface {
	Create(ctx context.Context, request web.UserHotelCreateRequest) web.UserHotelResponse
	Update(ctx context.Context, request web.UserHotelUpdateRequest) web.UserHotelResponse
	Delete(ctx context.Context, userHotelId int)
	FindById(ctx context.Context, userHotelId int) web.UserHotelResponse
	FindAll(ctx context.Context) []web.UserHotelResponse
}
