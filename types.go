package auth

type userCookie struct {
	userID        string
	tokenLocation string
}

// Authorization ...
type Authorization struct {
	Authorized    bool   `json:"authorized"`
	UserID        string `json:"userId"`
	TokenLocation string `json:"tokenLocation"`
}

// Token ...
type Token struct {
	Evenson   bool `json:"evenson"`
	Woodard   bool `json:"woodard"`
	Sundstedt bool `json:"sundstedt"`
}

// PassSet ...
type PassSet struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserInfo ...
type UserInfo struct {
	UserID string `json:"userId"`
	Token  Token  `json:"token"`
}
