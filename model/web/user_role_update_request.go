package web

type UserRoleUpdateRequest struct {
	Id            int `validate:"required"`
	RoleId        int `validate:"required" json:"role_id"`
	UserProfileId int `validate:"required" json:"user_profile_id"`
}
