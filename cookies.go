package auth

import (
	// std lib
	"io/ioutil"
	"log"
	"net/http"
	"time"

	// external lib
	"github.com/gorilla/securecookie"
)

const cookieName string = "Sundstedts-IAM"

var hashKey []byte

// SetCookie creates a secured http cookie that stores the location of the user's claims token
func SetCookie(w http.ResponseWriter, host, userID, tokenLocation string) error {
	h, err := getHashKey(host)
	if err != nil {
		return err
	}
	s := securecookie.New(*h, nil)

	v := map[string]string{
		"userID":        userID,
		"tokenLocation": tokenLocation,
	}

	encoded, err := s.Encode(cookieName, v)
	if err == nil {
		cookie := &http.Cookie{
			Name:     cookieName,
			Value:    encoded,
			Domain:   ".sundstedt.us",
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}

	return nil
}

// UnsetCookie removes the cookie for this user
func UnsetCookie(w http.ResponseWriter) {
	c := &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Domain:   ".sundstedt.us",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}

	http.SetCookie(w, c)
}

func getCookie(r *http.Request, host string) (*userCookie, error) {
	h, err := getHashKey(host)
	if err != nil {
		return nil, err
	}
	log.Println("HashKey:", h)

	s := securecookie.New(*h, nil)
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil, ErrNotFound
	}

	log.Println("Cookie", cookie)

	val := make(map[string]string)
	err = s.Decode(cookieName, cookie.Value, &val)
	if err != nil {
		return nil, err
	}

	u := userCookie{
		userID:        val["userID"],
		tokenLocation: val["tokenLocation"],
	}

	return &u, nil
}

func getHashKey(host string) (*[]byte, error) {
	if hashKey != nil {
		return &hashKey, nil
	}
	resp, err := http.Get(host + "/auth/cookiekey")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	hashKey = body

	return &hashKey, nil
}
