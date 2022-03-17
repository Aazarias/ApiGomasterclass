package model

type Games struct {
	ID          string `json:"id"`
	Name        string `json:"games_name"`
	Mark        string `json:"mark"`
	Price       string `json:"price"`
	RealaseDate string `json:"relase_date"`
}
