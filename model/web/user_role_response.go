package web

type UserRoleResponse struct {
	Id            int `json:"id"`
	RoleId        int `json:"role_id"`
	UserProfileId int `json:"user_profile_id"`
}
