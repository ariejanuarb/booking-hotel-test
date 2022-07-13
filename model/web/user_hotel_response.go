package web

type UserHotelResponse struct {
	Id            int `json:"id"`
	HotelId       int `json:"hotel_id"`
	UserProfileId int `json:"user_profile_id"`
}
