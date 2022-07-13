package web

type UserRoleCreateRequest struct {
	RoleId        int `json:"role_id"`
	UserProfileId int `json:"user_profile_id"`
}
