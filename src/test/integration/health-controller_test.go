package test_test

import (
	"encoding/json"
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/RaphaSalomao/gin-budget-control/config"
	"github.com/RaphaSalomao/gin-budget-control/database"
	"github.com/RaphaSalomao/gin-budget-control/router"
	"github.com/golang-migrate/migrate/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type HealthControllerSuite struct {
	suite.Suite
	db   *gorm.DB
	m    *migrate.Migrate
	port string
}

func (s *HealthControllerSuite) SetupSuite() {
	config.LoadEnvironment("../../test.env")
	s.Require().NoError(database.Connect())
	s.db = database.DB
	s.m = database.M
	s.port = router.SrvPort
	go router.RunServer()
	time.Sleep(2 * time.Second)
}

func (s *HealthControllerSuite) TearDownTest() {}

func (s *HealthControllerSuite) TearDownSuite() {
	s.m.Down()
	s.db.Exec("DROP TABLE schema_migrations")
	p, _ := os.FindProcess(os.Getpid())
	p.Signal(syscall.SIGINT)
}

func TestHealthControllerSuite(t *testing.T) {
	suite.Run(t, new(HealthControllerSuite))
}
func (s *HealthControllerSuite) TestHealthCheck_Success() {
	resp, err := http.Get("http://localhost:5000/budget-control/api/v1/health")
	s.Require().NoError(err)
	defer resp.Body.Close()

	expect := struct{ Online bool }{true}
	var got struct{ Online bool }
	json.NewDecoder(resp.Body).Decode(&got)

	s.Require().Equal(http.StatusOK, resp.StatusCode)
	s.Require().Equal(expect, got)
}
