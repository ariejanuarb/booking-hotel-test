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

type FloorServiceImpl struct {
	FloorRepository repository.FloorRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewFloorService(floorRepository repository.FloorRepository, DB *sql.DB, validate *validator.Validate) FloorService {
	return &FloorServiceImpl{
		FloorRepository: floorRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *FloorServiceImpl) Create(ctx context.Context, request web.FloorCreateRequest) web.FloorResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	floor := domain.Floor{
		Number:  request.Number,
		HotelId: request.HotelId,
	}

	floor = service.FloorRepository.Save(ctx, tx, floor)

	return helper.ToFloorResponse(floor)
}

func (service *FloorServiceImpl) Update(ctx context.Context, request web.FloorUpdateRequest) web.FloorResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	floor, err := service.FloorRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	floor.Number = request.Number
	floor.HotelId = request.HotelId

	floor = service.FloorRepository.Update(ctx, tx, floor)

	return helper.ToFloorResponse(floor)
}

func (service *FloorServiceImpl) Delete(ctx context.Context, floorId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	floor, err := service.FloorRepository.FindById(ctx, tx, floorId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.FloorRepository.Delete(ctx, tx, floor)
}

func (service *FloorServiceImpl) FindById(ctx context.Context, floorId int) web.FloorResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	floor, err := service.FloorRepository.FindById(ctx, tx, floorId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToFloorResponse(floor)
}

func (service *FloorServiceImpl) FindAll(ctx context.Context) []web.FloorResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	floors := service.FloorRepository.FindAll(ctx, tx)

	return helper.ToFloorResponses(floors)
}
