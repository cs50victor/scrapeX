package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"os/exec"
	"encoding/json"
	"io/ioutil"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func fiveAM(c *fiber.Ctx) error{	

	raw, err := ioutil.ReadFile("./5am.json")

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

func shouldMakeRequest() bool{
	timeFile, er := os.Stat("./5am.json")

	if er != nil {
		log.Fatal(er)
	}

	lastModified := timeFile.ModTime()
	currentTime := time.Now()
	sixHoursAgo := currentTime.Add(-6 * time.Hour)
	makeNewRequest := sixHoursAgo.After(lastModified)


	fmt.Println("File was last modified: ",lastModified.Format("2006-01-02 3:4:5pm"))
	fmt.Println("Time [6 Hours Ago]: ",sixHoursAgo.Format("2006-01-02 3:4:5pm"))
	fmt.Println("Make new Request? ",makeNewRequest)

	return makeNewRequest
}

func pythonRequests(){

	cmd := exec.Command("./stockTrend")	
	err := cmd.Run()
	
	if err != nil {
		fmt.Printf("Python exec run error: %s\n",err)
	}else{
		fmt.Println("Python exec ran successfully! New Requests Complete")
	}
}

// MAIN GO-FIBER SERVER
func main() {
	
	newRequest := shouldMakeRequest()
	if (newRequest){
		pythonRequests()
	}
	fmt.Println("New Request Made? ",newRequest)
	

	app := fiber.New()
	app.Use(logger.New())

	// api end-Points
	app.Get("/",hello)
	app.Get("/5am",fiveAM)

	port := os.Getenv("PORT")
	if os.Getenv("PORT") == ""{
		port = "3000"
	}
    // handle server starting error
    log.Fatal(app.Listen(":" +port))
}