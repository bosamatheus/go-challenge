package docs

import (
	"mercafacil-challenge/api/models"
)

// NOTE: Types defined here are purely for documentation purposes.

type UserRequest struct {
	Name      string
	Cellphone string
}

// swagger:parameters createUser updateUser
type UserParamsWrapper struct {
	// Data to create a new user.
	// in:body
	Body UserRequest
}

// Data structure representing a single product
// swagger:response userResponse
type UserResponseWrapper struct {
	// Newly created product.
	// in:body
	Body models.User
}

// A list of users
// swagger:response usersResponse
type UsersResponseWrapper struct {
	// All current users.
	// in: body
	Body []models.User
}
