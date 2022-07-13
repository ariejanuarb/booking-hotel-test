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

type UserRoleServiceImpl struct {
	UserRoleRepository repository.UserRoleRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewUserRoleService(userRoleRepository repository.UserRoleRepository, DB *sql.DB, validate *validator.Validate) UserRoleService {
	return &UserRoleServiceImpl{
		UserRoleRepository: userRoleRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *UserRoleServiceImpl) Create(ctx context.Context, request web.UserRoleCreateRequest) web.UserRoleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userRole := domain.UserRole{
		RoleId:        request.RoleId,
		UserProfileId: request.UserProfileId,
	}

	userRole = service.UserRoleRepository.Save(ctx, tx, userRole)

	return helper.ToUserRoleResponse(userRole)
}

func (service *UserRoleServiceImpl) Update(ctx context.Context, request web.UserRoleUpdateRequest) web.UserRoleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userRole, err := service.UserRoleRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	userRole.RoleId = request.RoleId
	userRole.UserProfileId = request.UserProfileId

	userRole = service.UserRoleRepository.Update(ctx, tx, userRole)

	return helper.ToUserRoleResponse(userRole)
}

func (service *UserRoleServiceImpl) Delete(ctx context.Context, userRoleId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userRole, err := service.UserRoleRepository.FindById(ctx, tx, userRoleId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRoleRepository.Delete(ctx, tx, userRole)
}

func (service *UserRoleServiceImpl) FindById(ctx context.Context, userRoleId int) web.UserRoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userRole, err := service.UserRoleRepository.FindById(ctx, tx, userRoleId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserRoleResponse(userRole)
}

func (service *UserRoleServiceImpl) FindAll(ctx context.Context) []web.UserRoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userRoles := service.UserRoleRepository.FindAll(ctx, tx)

	return helper.ToUserRoleResponses(userRoles)
}
