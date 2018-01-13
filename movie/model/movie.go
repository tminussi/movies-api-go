package movie

type Movie struct {
	Id    int64 `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Genre string `json:"genre" db:"genre"`
}
