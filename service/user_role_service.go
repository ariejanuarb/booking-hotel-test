package service

import (
	"context"
	"programmerzamannow/belajar-golang-restful-api/model/web"
)

type UserRoleService interface {
	Create(ctx context.Context, request web.UserRoleCreateRequest) web.UserRoleResponse
	Update(ctx context.Context, request web.UserRoleUpdateRequest) web.UserRoleResponse
	Delete(ctx context.Context, userRoleId int)
	FindById(ctx context.Context, userRoleId int) web.UserRoleResponse
	FindAll(ctx context.Context) []web.UserRoleResponse
}
