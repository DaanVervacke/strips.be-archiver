package types

type Series struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	CoverImage string  `json:"coverImage"`
	Intro      string  `json:"intro"`
	Albums     []Album `json:"albums"`
}
