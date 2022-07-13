package web

type UserHotelCreateRequest struct {
	HotelId       int `json:"hotel_id"`
	UserProfileId int `json:"user_profile_id"`
}
