package auth

import (
	// std lib
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	// external lib
	"github.com/gorilla/securecookie"
)

const ( 
	cookieName string = "Sundstedts-IAM"

	contentTypeAppJSON = "application/json"
)

type Authorization struct {
	Evenson bool `json:"evenson"`
	Woodard bool `json:"woodard"`
	Sundstedt bool `json:"sundstedt"`
}

// RegisterUser registers a new user on the sundstedts site (duh)
func RegisterUser(host, username, password string) (*Authorization, error) {
	fmt.Printf("Registering User %s, %s", username, password)
	requestBody, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(host + "/auth/register", contentTypeAppJSON, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%v", body)

	return nil, nil
}

// SetCookie creates a secured http cookie that stores the location of the user's claims token
func SetCookie(w http.ResponseWriter, hashKey []byte, userID, tokenLocation string) {
	s := securecookie.New(hashKey, nil)

	v := map[string]string {
		"userID": userID,
		"tokenLocation": tokenLocation,
	}

	encoded, err := s.Encode(cookieName, v)
	if err == nil {
		cookie := &http.Cookie{
			Name: cookieName,
			Value: encoded,
			Domain: ".sundstedt.us",
			Path: "/",
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
}

// UnsetCookie removes the cookie for this user
func UnsetCookie(w http.ResponseWriter) {
	c := &http.Cookie{
		Name: cookieName,
		Value: "",
		Domain: ".sundstedt.us",
		Path: "/",
		Expires: time.Unix(0, 0),
		HttpOnly: true,
	}

	http.SetCookie(w, c)
}

func main() {
	host := "http://localhost:7330"

	_, err := RegisterUser(host, "evan", "woodard")
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
}