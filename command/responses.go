package command

type changeNameResponse struct {
	Err string `json:"err"`
}

type createResponse struct {
	Id string `json:"id"`
	Err string `json:"err"`
}

type deleteResponse struct {
	Err string `json:"err"`
}

type turnOffResponse struct {
	Err string `json:"err"`
}

type turnOnResponse struct {
	Err string `json:"err"`
}