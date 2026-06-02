package internal

type Call struct {
	ID        int    `json:"id"`
	PhoneFrom string `json:"phone_from"`
	PhoneTo   string `json:"phone_to"`
	Country   string `json:"country"`
	Duration  int    `json:"duration"`
	Status    string `json:"status"`
}
