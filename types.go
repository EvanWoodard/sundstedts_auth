package auth

type userCookie struct {
	userID string
	tokenLocation string
}

type Authorization struct {
	Authorized bool `json:"authorized"`
	UserID string `json:"userId"`
	TokenLocation bool `json:"tokenLocation"`
}

type Token struct {
	Evenson bool `json:"evenson"`
	Woodard bool `json:"woodard"`
	Sundstedt bool `json:"sundstedt"`
}

type PassSet struct {
	Username string `json:"username"`
	Password string `json:"password"`
}