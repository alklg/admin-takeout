package sms

type User struct {
	AccountSid string `json:"accountSid"`
	AuthToken  string `json:"autoToken"`
	FromPhone  string `json:"fromPhone"`
}
