package factory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RaphaSalomao/gin-budget-control/model/entity"
	"github.com/RaphaSalomao/gin-budget-control/security"
	"gorm.io/gorm"
)

type Request struct {
	Method string
	Path   string
	Body   interface{}
	User   entity.User
	DB     *gorm.DB
	Client http.Client
	Port   string
}

func (r *Request) DoRequest() (*http.Response, error) {
	token, err := security.GenerateJWT(r.User.Email)
	if err != nil {
		return nil, err
	}
	tokenString := fmt.Sprintf("Bearer %s", token)

	requestBytes, err := json.Marshal(r.Body)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(requestBytes)

	url := fmt.Sprintf("http://localhost:%s%s", r.Port, r.Path)
	req, err := http.NewRequest(r.Method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", tokenString)

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *Request) SaveUser() {
	r.DB.Create(&r.User)
}
