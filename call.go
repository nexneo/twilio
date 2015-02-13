package twilio

type Call struct {
	AccountSid  string    `json:"account_sid"`
	DateCreated Timestamp `json:"date_created,omitempty"`
	DateUpdated Timestamp `json:"date_updated,omitempty"`
	From        string    `json:"from"`
	To          string    `json:"to"`
	Price       Price     `json:"price,omitempty"`
	Status      string    `json:"status"`
	Direction   string    `json:"direction"`
	Sid         string    `json:"sid"`
	Uri         string    `json:"uri"`
	AnsweredBy  string    `json:"answered_by"`
}
