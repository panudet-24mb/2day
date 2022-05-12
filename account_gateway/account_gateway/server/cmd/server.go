package cmd

import (
	"fmt"
	"identity_provider/internal/config"
	"identity_provider/internal/handler"
	"identity_provider/internal/middleware"
	"identity_provider/internal/repository"
	"identity_provider/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func Execute() {
	config.InitTimeZone()
	config.InitConfig()
	// ka := client.InitKafka()
	// fmt.Println(ka)
	// ka.EmitSync("some-key", "some-value")
	// ka.Finish()
	db := config.InitDatabase()
	db.AutoMigrate(&repository.User{}, &repository.Account{}, &repository.Address{},
		&repository.AddressGroup{}, &repository.AddressType{})
	userRepositoryDB := repository.NewUserRepositoryDB(db)
	accountRepositoryDB := repository.NewAccountRepositoryDB(db)
	userService := service.NewUserService(userRepositoryDB)
	authService := service.NewAuthService(userRepositoryDB)
	accountService := service.NewAccountService(accountRepositoryDB)
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)
	accountHandler := handler.NewAccountHandler(accountService)

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "IdentityProvider API V1",
	})
	app.Use(recover.New())
	app.Use(cors.New())
	api := app.Group("/api")
	v1 := api.Group("/v1")
	auth := api.Group("/auth")
	v1.Get("/users", userHandler.GetUsers)
	v1.Post("/users", userHandler.CreateUser)
	v1.Post("/users/accept-term-condition", middleware.Protected(), userHandler.AcceptTermCondition)
	v1.Get("/account", middleware.Protected(), accountHandler.FindAccount)
	v1.Post("/account/create-account", middleware.Protected(), accountHandler.CreateAccount)
	v1.Put("/account", middleware.Protected(), accountHandler.UpdateAccount)

	auth.Post("/login", authHandler.DefaultLogin)

	app.Get("/", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ hello")
		return c.SendString(msg) // => ✋ register
	})

	port := viper.GetString("app.port")
	app.Listen(":" + port)

}
