package tests

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	my_app "github.com/giftxtrade/api/src/app"
	"github.com/giftxtrade/api/src/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	server *fiber.App
	app *my_app.AppBase
	controller controllers.Controller
)

func NewMockDB() {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not ping pool: %s", err)
	}

	// Build and run the given Dockerfile
	resource, err := pool.RunWithOptions(
		&dockertest.RunOptions{
			Repository: "postgres",
			Tag:        "14",
			Env: []string{
				"POSTGRES_PASSWORD=postgres",
				"POSTGRES_USER=postgres",
				"POSTGRES_DB=postgres",
				"listen_addresses = '*'",
			},
		},
		func(config *docker.HostConfig) {
			// set AutoRemove to true so that stopped container goes away by itself
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{Name: "no"}
		},
	)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		hostAndPort := resource.GetHostPort("5432/tcp")
		db, err = sql.Open("postgres", fmt.Sprintf("postgres://postgres:postgres@%s/postgres?sslmode=disable", hostAndPort))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
}

func SetupMockController(app *my_app.AppBase) controllers.Controller {
	return controllers.Controller{
		AppContext: app.AppContext,
		Service: app.Service,
		Querier: app.Querier,
	}
}

func TestMain(m *testing.M) {
	NewMockDB()
	app = my_app.NewMock(db, fiber.New())
	server = fiber.New()
	controller = SetupMockController(app)

	// run tests...
	exitCode := m.Run()

	os.Exit(exitCode)
}
