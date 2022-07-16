package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"programmerzamannow/belajar-golang-restful-api/exception"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"programmerzamannow/belajar-golang-restful-api/model/domain"
	"programmerzamannow/belajar-golang-restful-api/model/web"
	"programmerzamannow/belajar-golang-restful-api/repository"
)

type UserHotelServiceImpl struct {
	UserHotelRepository repository.UserHotelRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewUserHotelService(userHotelRepository repository.UserHotelRepository, DB *sql.DB, validate *validator.Validate) UserHotelService {
	return &UserHotelServiceImpl{
		UserHotelRepository: userHotelRepository,
		DB:                  DB,
		Validate:            validate,
	}
}

func (service *UserHotelServiceImpl) Create(ctx context.Context, request web.UserHotelCreateRequest) web.UserHotelResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userHotel := domain.UserHotel{
		HotelId:       request.HotelId,
		UserProfileId: request.UserProfileId,
	}

	userHotel = service.UserHotelRepository.Save(ctx, tx, userHotel)

	return helper.ToUserHotelResponse(userHotel)
}

func (service *UserHotelServiceImpl) Update(ctx context.Context, request web.UserHotelUpdateRequest) web.UserHotelResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userHotel, err := service.UserHotelRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	userHotel.HotelId = request.HotelId

	userHotel = service.UserHotelRepository.Update(ctx, tx, userHotel)

	return helper.ToUserHotelResponse(userHotel)
}

func (service *UserHotelServiceImpl) Delete(ctx context.Context, userHotelId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userHotel, err := service.UserHotelRepository.FindById(ctx, tx, userHotelId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserHotelRepository.Delete(ctx, tx, userHotel)
}

func (service *UserHotelServiceImpl) FindById(ctx context.Context, userHotelId int) web.UserHotelResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userHotel, err := service.UserHotelRepository.FindById(ctx, tx, userHotelId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserHotelResponse(userHotel)
}

func (service *UserHotelServiceImpl) FindAll(ctx context.Context) []web.UserHotelResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userHotels := service.UserHotelRepository.FindAll(ctx, tx)

	return helper.ToUserHotelResponses(userHotels)
}
