// Contains struct for requests
package dto

type CreateTodoRequest struct {
	Title string 			`json:"title" binding:"required"`
	Completed bool 			`json:"completed"`
}