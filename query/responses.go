package query

type getAllOnResponse struct {
	AllOn []string `json:"allOn"`
	Err string `json:"err"`
}

type isOnResponse struct {
	IsOn bool `json:"isOn"`
	Err string `json:"err"`
}
