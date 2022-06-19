package verifyuser

type GEtheader struct {
	Jwt string `json:"jwt"`
	OTP string `json:"otp"`
}
type DATA struct {
	Email                       string
	Username, Tag, UserId, Time string

	Subdata struct {
		Password string
	}
}
