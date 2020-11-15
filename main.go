package auth

import (
	// std lib
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	contentTypeAppJSON = "application/json"
)

var (
	ErrNotFound = errors.New("Not Found")
)

// RegisterUser registers a new user on the sundstedts site (duh)
func RegisterUser(host, username, password string) (*Authorization, error) {
	requestBody, err := json.Marshal(PassSet{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(host+"/auth/register", contentTypeAppJSON, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	a := Authorization{}
	json.Unmarshal(body, &a)

	return &a, nil
}

// Login ...
func Login(host, username, password string) (*Authorization, error) {
	requestBody, err := json.Marshal(PassSet{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(host+"/auth/login", contentTypeAppJSON, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	a := Authorization{}
	json.Unmarshal(body, &a)

	return &a, nil
}

// GetToken ...
func GetToken(r *http.Request, host string) (*Token, error) {
	cookie, err := getCookie(r, host)
	if err != nil {
		return nil, err
	}

	tl := cookie.tokenLocation

	resp, err := http.Get(host + "/auth/token/" + tl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	t := Token{}
	json.Unmarshal(body, &t)

	return &t, nil
}

// GetUserInfo ...
func GetUserInfo(r *http.Request, host string) (*UserInfo, error) {
	var ui UserInfo

	cookie, err := getCookie(r, host)
	if err != nil {
		log.Println("Failed to get cookie", err)
		return nil, err
	}

	ui.UserID = cookie.userID

	tl := cookie.tokenLocation

	resp, err := http.Get(host + "/auth/token/" + tl)
	if err != nil {
		log.Println("failed to get token", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed to read response", err)
		return nil, err
	}

	t := Token{}
	json.Unmarshal(body, &t)

	ui.Token = t

	return &ui, nil
}
