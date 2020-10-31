package dto

type GuestLoginResponse struct {
	UserId         int64
	AccessToken    string
	NickName       string
	LoginName      string
	UserRole       string
	Country        string
	Channel        string
	UserChannel    string
	Avatar         string
	RegGuestId     string
	RegPkgName     string
	CreateTimeUnix int64
	AvailableCoin  int64
}
