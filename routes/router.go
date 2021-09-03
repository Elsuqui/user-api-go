package routes

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

// Define struct to declare all groups of routes divided by modules
type Router struct {
	UserR  UserRoutes
	LoginR LoginRoutes
}

// Function to initialize all api routes
func (r *Router) Boot(server *gin.Engine) {
	structValue := reflect.ValueOf(r)                  // Get reflection value of Router struct
	structElements := structValue.Elem()               // Get struct pointer
	inputs := []reflect.Value{reflect.ValueOf(server)} // Just input 1 parameter (server Engine pointer)

	for i := 0; i < structElements.NumField(); i++ {
		item := structElements.Field(i).Addr()             // Get address that is pointing this field
		item.MethodByName("InitializeRoutes").Call(inputs) // Call Initialize Routes method
	}
}
