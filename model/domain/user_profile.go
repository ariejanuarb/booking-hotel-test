package domain

type UserProfile struct {
	Id              int
	Email           string
	Password        string
	Name            string
	Gender          string
	AssignedToHotel bool
}
