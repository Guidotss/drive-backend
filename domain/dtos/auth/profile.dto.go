package auth

type ProfileDTO struct {
	Name     string `json:"name"`
	Bio      string `json:"bio"`
	Location string `json:"location"`
	Website  string `json:"website"`
	Avatar   string `json:"avatar"`
}
