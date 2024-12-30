package types

type Account struct {
	AccountStatus string    `json:"accountStatus"`
	Profiles      []Profile `json:"profiles"`
}

type Profile struct {
	ID string `json:"id"`
}
