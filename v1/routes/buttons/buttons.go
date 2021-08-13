package buttonsroutehandler

import (
	"fmt"
	"strings"
	fiber "github.com/gofiber/fiber/v2"
	try "github.com/manucorporat/try"
	utils "github.com/0187773933/FireStickC2Server/v1/utils"
)

var Routes map[string]func()(string)
var GlobalRedisConfig *utils.RedisConfig

func RegisterRoutes( fiber_app *fiber.App , redis_config *utils.RedisConfig ) {
	GlobalRedisConfig = redis_config
	Routes = map[string]func()(string) {
		"/pause": Pause ,
		"/resume": Resume ,
		"/play": Play ,
		"/stop": Stop ,
		"/next": Next ,
		"/previous": Previous ,
	}
	buttons := fiber_app.Group( "/buttons" )
	for key , _ := range Routes {
		buttons.Get( key , HandleFiberContext )
	}
}

func HandleFiberContext( context *fiber.Ctx ) ( error ) {
	path := context.Path()
	fmt.Println( path )
	function_name := strings.Split( path , "/buttons" )
	if len( function_name ) < 2 { return context.JSON( fiber.Map{ "result": "failed" } ) }
	_ , exists := Routes[ function_name[ 1 ] ]
	if exists == false { return context.JSON( fiber.Map{ "result": "failed" } ) }
	result := "failed"
	try.This( func() {
		result = Routes[ function_name[ 1 ] ]()
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON( fiber.Map{
		"route": path ,
		"result": result ,
	})
}

func Pause() ( result string ) {
	result = "success"
	fmt.Println( "pause" )
	return
}

func Resume() ( result string ) {
	result = "success"
	fmt.Println( "resume" )
	return
}

func Play() ( result string ) {
	result = "success"
	fmt.Println( "play" )
	fmt.Println( GlobalRedisConfig )
	return
}

func Stop() ( result string ) {
	result = "success"
	fmt.Println( "stop" )
	return
}

func Next() ( result string ) {
	result = "success"
	fmt.Println( "next" )
	return
}

func Previous() ( result string ) {
	result = "success"
	fmt.Println( "previous" )
	return
}