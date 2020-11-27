package command

type changeNameRequest struct {
	Id string `json:"id"`
	NewName string `json:"newName"`
}

type createRequest struct {
	Name string `json:"name"`
}

type deleteRequest struct {
	Id string `json:"id"`
}

type turnOffRequest struct {
	Id string `json:"id"`
}

type turnOnRequest struct {
	Id string `json:"id"`
}