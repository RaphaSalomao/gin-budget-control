package test_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/RaphaSalomao/gin-budget-control/database"
	"github.com/RaphaSalomao/gin-budget-control/model"
	"github.com/RaphaSalomao/gin-budget-control/model/entity"
	"github.com/RaphaSalomao/gin-budget-control/model/enum"
	"github.com/RaphaSalomao/gin-budget-control/router"
	"github.com/RaphaSalomao/gin-budget-control/test/factory"
	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ControllerSuite struct {
	suite.Suite
	db   *gorm.DB
	m    *migrate.Migrate
	port string
}

func (s *ControllerSuite) SetupSuite() {
	s.Require().NoError(godotenv.Load("../../test.env"))
	s.Require().NoError(database.Connect())
	s.db = database.DB
	s.m = database.M
	s.port = os.Getenv("SRV_PORT")
	go router.HandleRequests()
	time.Sleep(2 * time.Second)
}

func (s *ControllerSuite) TearDownTest() {
	s.db.Exec("DELETE FROM receipts")
	s.db.Exec("DELETE FROM expenses")
	s.db.Exec("DELETE FROM users")
}

func (s *ControllerSuite) TearDownSuite() {
	s.m.Down()
	s.db.Exec("DROP TABLE schema_migrations")
	p, _ := os.FindProcess(os.Getpid())
	p.Signal(syscall.SIGINT)
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(ControllerSuite))
}

func (s *ControllerSuite) TestMonthBalanceSumary_Success() {
	// prepare database
	r := factory.Request{
		User: entity.User{
			Email:    "email@email.com",
			Password: "password",
		},
		Method: http.MethodGet,
		DB:     s.db,
		Client: http.Client{},
		Port:   s.port,
	}
	r.SaveUser()

	receipts := []entity.Receipt{
		{
			Description: "Receipt 1",
			Value:       1100,
			Date:        "2020-01-01T00:00:00Z",
			UserId:      r.User.Id,
		},
		{
			Description: "Receipt 2",
			Value:       1200,
			Date:        "2020-01-02T00:00:00Z",
			UserId:      r.User.Id,
		},
		{
			Description: "Receipt 3",
			Value:       1300,
			Date:        "2020-01-03T00:00:00Z",
			UserId:      r.User.Id,
		},
	}
	expenses := []entity.Expense{
		{
			Description: "Expense 1",
			Value:       1100,
			Date:        "2020-01-01T00:00:00Z",
			Category:    enum.CategoryFood,
			UserId:      r.User.Id,
		},
		{
			Description: "Expense 2",
			Value:       250,
			Date:        "2020-01-02T00:00:00Z",
			Category:    enum.CategoryHealth,
			UserId:      r.User.Id,
		},
		{
			Description: "Expense 3",
			Value:       100,
			Date:        "2020-01-03T00:00:00Z",
			Category:    enum.CategoryHealth,
			UserId:      r.User.Id,
		},
		{
			Description: "Expense 4",
			Value:       1000,
			Date:        "2020-01-04T00:00:00Z",
			Category:    enum.CategoryFood,
			UserId:      r.User.Id,
		},
	}
	s.db.Create(&receipts)
	s.db.Create(&expenses)

	// prepare expected response
	totalReceipt := 0.0
	totalExpense := 0.0
	categoryBalance := map[enum.Category]float64{}

	for _, receipt := range receipts {
		totalReceipt += receipt.Value
	}
	for _, expense := range expenses {
		totalExpense += expense.Value
		categoryBalance[expense.Category] += expense.Value
	}
	monthBalance := totalReceipt - totalExpense

	// do request
	year, month := "2020", "01"
	r.Path = fmt.Sprintf("/budget-control/api/v1/summary/%s/%s", year, month)
	resp, err := r.DoRequest()
	s.Require().NoError(err)

	// assert response
	var bs model.BalanceSumaryResponse
	json.NewDecoder(resp.Body).Decode(&bs)

	s.Require().Equal(http.StatusOK, resp.StatusCode)
	s.Require().Equal(monthBalance, bs.MonthBalance)
	s.Require().Equal(totalReceipt, bs.TotalReceipt)
	s.Require().Equal(totalExpense, bs.TotalExpense)
	s.Require().Equal(categoryBalance, bs.CategoryBalance)
}
