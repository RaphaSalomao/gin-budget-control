package test_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/RaphaSalomao/gin-budget-control/database"
	"github.com/RaphaSalomao/gin-budget-control/model/entity"
	"github.com/RaphaSalomao/gin-budget-control/router"
	"github.com/RaphaSalomao/gin-budget-control/security"
	"github.com/RaphaSalomao/gin-budget-control/utils"
	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserControllerSuite struct {
	suite.Suite
	db   *gorm.DB
	m    *migrate.Migrate
	port string
}

func (s *UserControllerSuite) SetupSuite() {
	s.Require().NoError(godotenv.Load("../../test.env"))
	s.Require().NoError(database.Connect())
	s.db = database.DB
	s.m = database.M
	s.port = os.Getenv("SRV_PORT")
	go router.HandleRequests()
	time.Sleep(2 * time.Second)
}

func (s *UserControllerSuite) TearDownTest() {
	s.db.Exec("DELETE FROM users")
}

func (s *UserControllerSuite) TearDownSuite() {
	s.m.Down()
	s.db.Exec("DROP TABLE schema_migrations")
	p, _ := os.FindProcess(os.Getpid())
	p.Signal(syscall.SIGINT)
}

func TestUserControllerSuite(t *testing.T) {
	suite.Run(t, new(UserControllerSuite))
}

func (s *UserControllerSuite) TestCreateUser_Success() {
	// prepare request
	expect := entity.UserRequest{
		Email:    "email@email.com",
		Password: "password",
	}

	request, err := json.Marshal(expect)
	s.Require().NoError(err)
	requestBody := bytes.NewBuffer(request)

	// do request
	resp, err := http.Post("http://localhost:5000/budget-control/api/v1/user", "application/json", requestBody)

	// assert response
	s.Require().NoError(err)
	s.Require().Equal(http.StatusCreated, resp.StatusCode)

	var user entity.User
	s.db.Where("email = ?", expect.Email).First(&user)
	s.Require().Equal(expect.Email, user.Email)
	s.Require().Equal(true, utils.ValidadePasswordHash(expect.Password, user.Password))
}

func (s *UserControllerSuite) TestAuthenticate_Success() {
	// prepare database
	password := utils.HashPassword("password")
	user := entity.User{
		Email:    "email@email.com",
		Password: password,
	}
	s.db.Create(&user)

	// prepare request
	userRequest := entity.UserRequest{
		Email:    "email@email.com",
		Password: "password",
	}

	request, err := json.Marshal(userRequest)
	s.Require().NoError(err)
	requestBody := bytes.NewBuffer(request)

	// do request
	resp, err := http.Post("http://localhost:5000/budget-control/api/v1/user/authenticate", "application/json", requestBody)
	s.Require().NoError(err)

	var tokenResponse struct{ Token string }
	json.NewDecoder(resp.Body).Decode(&tokenResponse)
	loggedUser, err := security.GetLoggedUserFromToken(tokenResponse.Token)

	// assert response
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.StatusCode)
	s.Require().NoError(err)
	s.Require().Equal(user.Id, loggedUser.Id)
}

func (s *UserControllerSuite) TestAuthenticate_Fail() {
	// prepare request
	userRequest := entity.UserRequest{
		Email:    "email@email.com",
		Password: "password",
	}

	request, err := json.Marshal(userRequest)
	s.Require().NoError(err)
	requestBody := bytes.NewBuffer(request)

	// do request
	resp, err := http.Post("http://localhost:5000/budget-control/api/v1/user/authenticate", "application/json", requestBody)

	var tokenResponse struct{ Token string }
	json.NewDecoder(resp.Body).Decode(&tokenResponse)

	// assert response
	s.Require().NoError(err)
	s.Require().Equal(http.StatusUnauthorized, resp.StatusCode)
}
