package handlers

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

// GetSwaggerSpec - Returns an OpenAPI specification file
func GetSwaggerSpec(ctx *fasthttp.RequestCtx) {
	log.Println("request for swagger specs")

	// TODO: Return an OpenAPI specification file

	fmt.Fprintf(ctx, "You have requested for an OpenAPI specification\n")
}

// GetAllInstances - Returns all avaliable instances
func GetAllInstances(ctx *fasthttp.RequestCtx) {
	log.Println("request for all instances")

	// TODO: Return user instances array if present, Not Found error otherwise

	fmt.Fprintf(ctx, "You have requested for all instances\n")
}

// GetInstancesByID - Returns information about instance by ID
func GetInstancesByID(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	log.Println("request for instance with id", id)

	// TODO: Return user's instance by id

	fmt.Fprintf(ctx, "You have requested for instance with id \"%v\"\n", id)
}

// DeleteInstanceByID - Deletes an instance from list
func DeleteInstanceByID(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	log.Println("request to delete instance with id", id)

	// TODO: Delete user's instance by id

	fmt.Fprintf(ctx, "You have requested for deletion of instance with id \"%v\"\n", id)
}

// GetInstanceHistoryByID - Returns history timeline by ID
func GetInstanceHistoryByID(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	log.Println("request to get history for instance with id", id)

	// TODO: Implement instance history

	fmt.Fprintf(ctx, "You have requested for history for instance with id \"%v\"\n", id)
}
