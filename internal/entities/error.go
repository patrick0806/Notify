package entities

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
