package service

import (
	"context"
	"programmerzamannow/belajar-golang-restful-api/model/web"
)

type UserProfileService interface {
	Create(ctx context.Context, request web.UserProfileCreateRequest) web.UserProfileResponse
	Update(ctx context.Context, request web.UserProfileUpdateRequest) web.UserProfileResponse
	Delete(ctx context.Context, userProfileId int)
	FindById(ctx context.Context, userProfileId int) web.UserProfileResponse
	FindAll(ctx context.Context) []web.UserProfileResponse
}
