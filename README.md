
![Logo](https://repository-images.githubusercontent.com/234231371/00fd8700-5430-11ea-820b-15fd85b2472c)


# Golang-Fiber

Fiber is an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for fast development with zero memory allocation and performance in mind.




## Hello Fiber (Golang Framework)

```go
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
```



## Running app

Run this command to start your server

```bash
  go run main.go
```
    
## Result
Browse to http://localhost:3000 and you should see Hello, Fiber Message! on the page.

