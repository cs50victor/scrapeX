package main

import (
	"fmt"
	"log"
	"os/exec"
	"io/ioutil"
	"encoding/json"
    "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	
	cmd := exec.Command("python","-c","from stockTrend import getNewJson; print(getNewJson())")	
	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("python file execution error: %s",err.Error())
	}else{
		fmt.Printf("python file output \n\n%s",output)
	}
	
	app := fiber.New()
	app.Use(logger.New())


	app.Get("/",hello)

	app.Get("/stocks", stocks)
	app.Get("/safe-crypto", safeCypto)
	app.Get("/risky-crypto", riskyCypto)

	serverErr := app.Listen(":3000")

    // handle server starting error
    if serverErr != nil {
        panic(serverErr)
    }
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func stocks(c *fiber.Ctx) error{	

	raw, err := ioutil.ReadFile("./5am/res/stocks.json")

	if err != nil {
		fmt.Println("opening Json error",err)
	}else{
		fmt.Println("Successfully Opened json")
	}
	

	var jsonFile map[string]interface{}
	err = json.Unmarshal(raw, &jsonFile)
	if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }

	return c.JSON(jsonFile)
}

func safeCypto(c *fiber.Ctx) error{	

	raw, err := ioutil.ReadFile("./5am/res/safe-crypto.json")

	if err != nil {
		fmt.Println("opening Json error",err)
	}else{
		fmt.Println("Successfully Opened json")
	}
	

	var jsonFile map[string]interface{}
	err = json.Unmarshal(raw, &jsonFile)
	if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }

	return c.JSON(jsonFile)
}

func riskyCypto(c *fiber.Ctx) error{	

	raw, err := ioutil.ReadFile("./5am/res/risky-crypto.json")

	if err != nil {
		fmt.Println("opening Json error",err)
	}else{
		fmt.Println("Successfully Opened json")
	}
	

	var jsonFile map[string]interface{}
	err = json.Unmarshal(raw, &jsonFile)
	if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }

	return c.JSON(jsonFile)
}