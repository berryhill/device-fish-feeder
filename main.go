package main

import (
	"fmt"

	"github.com/device-fish-feeder/models"
	"github.com/device-fish-feeder/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

const DeviceId = "1234"

func main() {
  	fmt.Println("Starting Fish Feeder Device")
	e := echo.New()

	controllers.Device = models.NewDevice(DeviceId)
	controllers.StartMqttClient()

	fmt.Println("Running a Server on localhost:1323")
	e.Run(standard.New(":1323"))
}

