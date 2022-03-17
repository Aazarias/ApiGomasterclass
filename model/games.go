package model

//Games represent a real game with his name, mark, price and release date
type Games struct {
	ID          string `json:"id"`
	Name        string `json:"games_name"`
	Mark        string `json:"mark"`
	Price       string `json:"price"`
	RealaseDate string `json:"release_date"`
}
