package web

type UserProfileResponse struct {
	Id              int    `json:"id"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Name            string `json:"name"`
	Gender          string `json:"gender"`
	AssignedToHotel bool   `json:"assigned_to_hotel"`
}
