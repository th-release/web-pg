package auth

type GetKeyDto struct {
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}
