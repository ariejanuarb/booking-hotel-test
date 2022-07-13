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

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewRoleService(roleRepository repository.RoleRepository, DB *sql.DB, validate *validator.Validate) RoleService {
	return &RoleServiceImpl{
		RoleRepository: roleRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *RoleServiceImpl) Create(ctx context.Context, request web.RoleCreateRequest) web.RoleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role := domain.Role{
		Type: request.Type,
	}

	role = service.RoleRepository.Save(ctx, tx, role)

	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) Update(ctx context.Context, request web.RoleUpdateRequest) web.RoleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role, err := service.RoleRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	role.Type = request.Type

	role = service.RoleRepository.Update(ctx, tx, role)

	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) Delete(ctx context.Context, roleId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role, err := service.RoleRepository.FindById(ctx, tx, roleId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.RoleRepository.Delete(ctx, tx, role)
}

func (service *RoleServiceImpl) FindById(ctx context.Context, roleId int) web.RoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role, err := service.RoleRepository.FindById(ctx, tx, roleId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) FindAll(ctx context.Context) []web.RoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roles := service.RoleRepository.FindAll(ctx, tx)

	return helper.ToRoleResponses(roles)
}
