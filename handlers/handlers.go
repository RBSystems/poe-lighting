package handlers

import (
	"fmt"
	"net/http"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/poe-lighting/helpers"
	"github.com/labstack/echo"
)

//TurnOn- turns the designated light on
func TurnOn(context echo.Context) error {
	ipaddress := context.Param("address")
	light := context.Param("light")
	helpers.LightOn(ipaddress, light)
	log.L.Infof("%s", light)
	return context.JSON(http.StatusOK, fmt.Sprintf("light %s was turned on", light))
}

//Turnoff - Turns the designated light off
func TurnOff(context echo.Context) error {
	ipaddress := context.Param("address")
	light := context.Param("light")
	helpers.LightOff(ipaddress, light)
	log.L.Infof("%s", light)

	return context.JSON(http.StatusOK, fmt.Sprintf("light %s was turned off", light))
}

//DimLight- this function dims the light
func DimLight(context echo.Context) error {
	ipaddress := context.Param("address")
	light := context.Param("light")
	dimlevel := context.Param("dimlevel")
	helpers.Dim(ipaddress, light, dimlevel)
	log.L.Infof("%s", light)
	log.L.Infof("%s", dimlevel)
	return context.JSON(http.StatusOK, fmt.Sprintf("light %s was set to %s", light, dimlevel))
}
