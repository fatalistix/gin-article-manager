package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Default gin router
	router := gin.Default()

	// load HTML templates located in templates folder.
	// Once loaded, these don't have to be read again on every request
	// Making Gin web applications very fast (lmao)
	router.LoadHTMLGlob("./templates/*")

	// Heart of Gin - divide the applications into various routes and define handlers for each route.
	// Creation of a route for the index page and an inline route handler.
	// This method is used to define a route handler for a GET request. It takes in as parameters to route (/)
	// and one or more route handlers which are just functions.
	// The route handler has a pointer to the context (gin.Context) as its parameter.
	// This context contains all the information about the request that the handler might need to process it.
	// For example, it includes information about the headers, cookies, etc.
	// The Context also has methods to render a response in HTML, text, JSON and XML formats. In this case, we use the
	// context.HTML method to render an HTML template (index.html). The call to this method includes additional data in
	// which the value of `title` is set to `Home Page`. This is a value that the HTML template can make use of.
	// In this case, we use this value in the `<title>` tag in the header's template.
	router.GET("/", func(c *gin.Context) {
		// Call the HTML method of the gin.Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)
	})

	// Start application on `localhost` and serve on the 8080 port by default
	fmt.Println(router.Run())
}
