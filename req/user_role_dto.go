package req

type (
	// UserRoleDto 用户角色dto
	UserRoleDto struct {
		ID       uint64 `json:"id"`
		UserId   uint64 `json:"user_id"`
		RoleId   uint64 `json:"role_id"`
		RoleCode string `json:"role_code"`
	}
)
