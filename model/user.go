package model

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Password string

func (p *Password) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}

	h := sha256.New()
	h.Write([]byte(str))
	*p = Password(fmt.Sprintf("%x", h.Sum(nil)))

	return nil
}

func (p Password) MarshalJSON() ([]byte, error) {
	return json.Marshal("")
}

type LoginUser struct {
	Email    string    `json:"email"`
	Password *Password `json:"pass,omitempty"`
}

type User struct {
	Id        string    `json:"id"`
	Lastname  string    `json:"name"`
	Firstname string    `json:"firstname"`
	Email     string    `json:"email"`
	Age       uint8     `json:"age"`
	Password  *Password `json:"pass,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User
