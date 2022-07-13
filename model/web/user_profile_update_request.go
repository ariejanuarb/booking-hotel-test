package web

type UserProfileUpdateRequest struct {
	Id              int    `json:"id"`
	Email           string `validate:"required,min=1,max=100" json:"email"`
	Password        string `validate:"required,min=1,max=100" json:"password"`
	Name            string `validate:"required,min=1,max=100" json:"name"`
	Gender          string `validate:"required,min=1,max=100" json:"gender"`
	AssignedToHotel bool   `json:"assigned_to_hotel"`
}
