package main

import (
	"fmt"
	"log"
	"taskies/user"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	_ "github.com/lib/pq"
)

func main() {
	dsn := "host=localhost user=postgres password=admin dbname=taskies port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if err != nil {
		log.Fatal(err)
	}

	userRepository := user.NewUserRepository(db, redis)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        3,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World",
		})
	})
	app.Get("/users/:id", userHandler.GetByID)
	log.Fatal(app.Listen(fmt.Sprintf(":3000")))
}
