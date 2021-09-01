package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"os/exec"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// ENDPOINTS
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

func checkStatus(c *fiber.Ctx) error{
	timeFile, er := os.Stat("./5am.json")

	if er != nil {
		log.Fatal(er)
	}

	lastModified := timeFile.ModTime()
	currentTime := time.Now()
	sixHoursAgo := currentTime.Add(-6 * time.Hour)
	makeNewRequest := sixHoursAgo.After(lastModified)
	pythonExecStatus := pythonRequests()

	status := map[string]string{"PythonRequestStatus":pythonExecStatus,"File-last-modified": lastModified.Format("2006-01-02 3:4:5pm"),"Time_[6_Hours_Ago]": sixHoursAgo.Format("2006-01-02 3:4:5pm"),"Make_new_Request?": strconv.FormatBool(makeNewRequest)}
	
	return c.JSON(status)
}

// UTILS
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
	fmt.Println("Made new Request? ",makeNewRequest)

	return makeNewRequest
}

func pythonRequests()string {

	cmd := exec.Command("./stockTrend")	
	err := cmd.Run()
	runStatus := "..."
	
	if err != nil {
		runStatus = fmt.Sprintf("Python exec run error: %s", err)
	}else{
		runStatus = "Python exec ran successfully! New Requests Complete"
	}

	return runStatus
}

// MAIN GO-FIBER SERVER
func main() {
	
	newRequest := shouldMakeRequest()
	//if (newRequest){
	//	pythonRequests()
	//}
	pythonRequests()
	fmt.Println("New Request Made? ",newRequest)
	

	app := fiber.New()
	app.Use(logger.New())
	// Default config to allow-cross-origin 
	app.Use(cors.New())

	// api end-Points
	app.Get("/",hello)
	app.Get("/5am",fiveAM)
	app.Get("/5am/status", checkStatus)

	port := os.Getenv("PORT")
	if os.Getenv("PORT") == ""{
		port = "3000"
	}
    // handle server starting error
    log.Fatal(app.Listen(":" +port))
}