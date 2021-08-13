package twitchroutehandler

import (
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	try "github.com/manucorporat/try"
)

func Pause( context *fiber.Ctx ) ( error ) {
	result := "failed"
	try.This( func() {
		result = "gathered info here as json"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON( fiber.Map{
		"route": "/app/twitch/pause" ,
		"result": result ,
	})
}

func Resume( context *fiber.Ctx ) ( error ) {
	result := "failed"
	try.This( func() {
		result = "gathered info here as json"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON( fiber.Map{
		"route": "/button/resume" ,
		"result": result ,
	})
}

func Play( context *fiber.Ctx ) ( error ) {
	result := "failed"
	try.This( func() {
		result = "gathered info here as json"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON( fiber.Map{
		"route": "/app/twitch/play" ,
		"result": result ,
	})
}

func Stop( context *fiber.Ctx ) ( error ) {
	result := "failed"
	try.This( func() {
		result = "gathered info here as json"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON( fiber.Map{
		"route": "/app/twitch/stop" ,
		"result": result ,
	})
}

func Next( context *fiber.Ctx ) ( error ) {
	result := "failed"
	try.This( func() {
		result = "gathered info here as json"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON( fiber.Map{
		"route": "/app/twitch/next" ,
		"result": result ,
	})
}

func Previous( context *fiber.Ctx ) ( error ) {
	result := "failed"
	try.This( func() {
		result = "gathered info here as json"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON( fiber.Map{
		"route": "/app/twitch/previous" ,
		"result": result ,
	})
}