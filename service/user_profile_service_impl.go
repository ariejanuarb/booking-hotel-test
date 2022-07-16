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

type UserProfileServiceImpl struct {
	UserProfileRepository repository.UserProfileRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewUserProfileService(userProfileRepository repository.UserProfileRepository, DB *sql.DB, validate *validator.Validate) UserProfileService {
	return &UserProfileServiceImpl{
		UserProfileRepository: userProfileRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

func (service *UserProfileServiceImpl) Create(ctx context.Context, request web.UserProfileCreateRequest) web.UserProfileResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userProfile := domain.UserProfile{
		Name:     request.Name,
		Gender:   request.Gender,
		Email:    request.Email,
		Password: request.Password,
	}

	userProfile = service.UserProfileRepository.Save(ctx, tx, userProfile)

	return helper.ToUserProfileResponse(userProfile)
}

func (service *UserProfileServiceImpl) Update(ctx context.Context, request web.UserProfileUpdateRequest) web.UserProfileResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userProfile, err := service.UserProfileRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	userProfile.Email = request.Email
	userProfile.Password = request.Password
	userProfile.Name = request.Name
	userProfile.Gender = request.Gender
	userProfile.RoleId = request.RoleId

	userProfile = service.UserProfileRepository.Update(ctx, tx, userProfile)

	return helper.ToUserProfileResponse(userProfile)
}

func (service *UserProfileServiceImpl) Delete(ctx context.Context, userProfileId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userProfile, err := service.UserProfileRepository.FindById(ctx, tx, userProfileId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserProfileRepository.Delete(ctx, tx, userProfile)
}

func (service *UserProfileServiceImpl) FindById(ctx context.Context, userProfileId int) web.UserProfileResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userProfile, err := service.UserProfileRepository.FindById(ctx, tx, userProfileId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserProfileResponse(userProfile)
}

func (service *UserProfileServiceImpl) FindAll(ctx context.Context) []web.UserProfileResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userProfiles := service.UserProfileRepository.FindAll(ctx, tx)

	return helper.ToUserProfileResponses(userProfiles)
}
