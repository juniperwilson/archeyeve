package database

type Building struct {
	Lat     float64 `json:"lat" db:"lat"`
	Lng     float64 `json:"lng" db:"lng"`
	Address string `json:"address" db:"address"`
	Name string	`json:"name" db:"name"`
}
