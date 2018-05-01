package controllers

import (
	"errors"
	"net/http"
)

const inquirySecret = "7asf89asfd8s9addf"

type inquiry struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	secret  string
}

func (inq inquiry) validate() error {
	if inq.Email == "" {
		return errors.New("email required")
	}

	if inq.Name == "" {
		return errors.New("name required")
	}

	if inq.secret != inquirySecret {
		return errors.New("inquiry invalid")
	}

	return nil
}

// NewInquiry Request handler for creating a new inquiry
func NewInquiry(w http.ResponseWriter, r *http.Request) {
	inq := inquiry{
		Email:   r.FormValue("email"),
		Name:    r.FormValue("name"),
		Message: r.FormValue("message"),
		secret:  r.FormValue("secret"),
	}

	if reject(w, inq) {
		return
	}

	// TODO: save the inquiry into a database.

	emitAsJSON(w, inq)
}
