package helper

import (
	"programmerzamannow/belajar-golang-restful-api/model/domain"
	"programmerzamannow/belajar-golang-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToRoleResponse(role domain.Role) web.RoleResponse {
	return web.RoleResponse{
		Id:   role.Id,
		Type: role.Type,
	}
}

func ToRoleResponses(roles []domain.Role) []web.RoleResponse {
	var roleResponses []web.RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, ToRoleResponse(role))
	}
	return roleResponses
}

func ToUserRoleResponse(userRole domain.UserRole) web.UserRoleResponse {
	return web.UserRoleResponse{
		Id:            userRole.Id,
		RoleId:        userRole.RoleId,
		UserProfileId: userRole.UserProfileId,
	}
}

func ToUserRoleResponses(userRoles []domain.UserRole) []web.UserRoleResponse {
	var userRoleResponses []web.UserRoleResponse
	for _, userRole := range userRoles {
		userRoleResponses = append(userRoleResponses, ToUserRoleResponse(userRole))
	}
	return userRoleResponses
}

func ToUserProfileResponse(userProfile domain.UserProfile) web.UserProfileResponse {
	return web.UserProfileResponse{
		Id:              userProfile.Id,
		Email:           userProfile.Email,
		Password:        userProfile.Password,
		Name:            userProfile.Name,
		Gender:          userProfile.Gender,
		AssignedToHotel: userProfile.AssignedToHotel,
	}
}

func ToUserProfileResponses(userProfiles []domain.UserProfile) []web.UserProfileResponse {
	var userProfileResponses []web.UserProfileResponse
	for _, userProfile := range userProfiles {
		userProfileResponses = append(userProfileResponses, ToUserProfileResponse(userProfile))
	}
	return userProfileResponses
}

func ToUserHotelResponse(userHotel domain.UserHotel) web.UserHotelResponse {
	return web.UserHotelResponse{
		Id:            userHotel.Id,
		HotelId:       userHotel.HotelId,
		UserProfileId: userHotel.UserProfileId,
	}
}

func ToUserHotelResponses(userHotels []domain.UserHotel) []web.UserHotelResponse {
	var userHotelResponses []web.UserHotelResponse
	for _, userHotel := range userHotels {
		userHotelResponses = append(userHotelResponses, ToUserHotelResponse(userHotel))
	}
	return userHotelResponses
}

func ToHotelResponse(hotel domain.Hotel) web.HotelResponse {
	return web.HotelResponse{
		Id:       hotel.Id,
		Name:     hotel.Name,
		Address:  hotel.Address,
		Province: hotel.Province,
		City:     hotel.City,
		ZipCode:  hotel.ZipCode,
		Star:     hotel.Star,
	}
}

func ToHotelResponses(hotels []domain.Hotel) []web.HotelResponse {
	var hotelResponses []web.HotelResponse
	for _, hotel := range hotels {
		hotelResponses = append(hotelResponses, ToHotelResponse(hotel))
	}
	return hotelResponses
}

func ToFloorResponse(floor domain.Floor) web.FloorResponse {
	return web.FloorResponse{
		Id:      floor.Id,
		Number:  floor.Number,
		HotelId: floor.HotelId,
		RoomId:  floor.RoomId,
	}
}

func ToFloorResponses(floors []domain.Floor) []web.FloorResponse {
	var floorResponses []web.FloorResponse
	for _, floor := range floors {
		floorResponses = append(floorResponses, ToFloorResponse(floor))
	}
	return floorResponses
}
