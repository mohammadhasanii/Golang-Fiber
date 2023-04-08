package main

import "github.com/gofiber/fiber/v2"


func bootstrap(){
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber (Framework Golang)")
	  })
    
	  app.Listen(":3000")
    
}


func main(){

	bootstrap()
}