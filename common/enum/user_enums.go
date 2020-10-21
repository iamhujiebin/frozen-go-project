package enum

type UserRole string

type userRoleEnumT struct {
	Normal UserRole
	Anchor UserRole
}

var UserRoleEnum = userRoleEnumT{
	Normal: "normal",
	Anchor: "anchor",
}
