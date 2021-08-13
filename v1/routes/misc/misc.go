package miscroutehandler

import (
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	try "github.com/manucorporat/try"
)

func Info( context *fiber.Ctx ) ( error ) {
	result := "failed"
	try.This( func() {
		result = "gathered info here as json"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON( fiber.Map{
		"route": "/misc/info" ,
		"result": result ,
	})
}