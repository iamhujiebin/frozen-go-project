package enum

type commonParamsT struct {
	UserId   string
	GuestId  string
	Platform string
	Version  string
	Uch      string
	Pch      string
	PkgName  string
	Country  string
}

var CommonParams = commonParamsT{
	UserId:   "__user_id",
	GuestId:  "__guest_id",
	Platform: "__platform",
	Version:  "__version",
	Uch:      "__uch",
	Pch:      "__pch",
	PkgName:  "__pname",
	Country:  "__country",
}
