package enum

type UserRole string

type userRoleEnumT struct {
	Normal UserRole
}

var UserRoleEnum = userRoleEnumT{
	Normal: "normal",
}
