package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"scrapper/utils"
	// "github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// ENDPOINTS
func hey(c *fiber.Ctx) error{

	var user string = c.Query("user")
	fmt.Println("user: " + user)

	var data string 
	if (len(user)==0){
		data = utils.Hello("User")
	} else{
		data = utils.Hello(user) 
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data": data,
	})
}

func coinbase(c *fiber.Ctx) error{
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data": utils.CommonCoinbaseTkns(),
	})
}

// MAIN GO-FIBER SERVER
func main() {
	
	//* -------- SERVER SETUP ---------
	app := goFiberApiSetup()
	
	// api end-Points
	app.Get("/",hey)
	app.Get("/tokens",coinbase)

	var port string = os.Getenv("PORT")
	if port == ""{
		port = "3000"
	}
    // handle server starting error
    log.Fatal(app.Listen(":" + port))
}

func goFiberApiSetup() *fiber.App{
	// -------- GO-FIBER SERVER SETUP ---------

	app := fiber.New()
	// Default config to allow-cross-origin 
	app.Use(cors.New())
	app.Use(logger.New())

	// cache (sends back the same response within time period)
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration: 10*60 * time.Second, // (15 minutes)
		CacheControl: true,
	}))

	// max of 5 requests every 10 seconds 
	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:5,
	}))

	return app
}
