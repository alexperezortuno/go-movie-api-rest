package types

type Error struct {
	Code    int16  `json:"code"`
	Message string `json:"Message"`
}

type ResponseDetail struct {
	Code    int16
	Message string
}
