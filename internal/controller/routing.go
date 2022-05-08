package controller

import "net/http"

func (c *Controller) routing() {
	// swagger:route GET /user/{userId} Users getUsers
	// Returns a user for a given ID.
	//
	// Parameters:
	//   + name: userId
	//     in: path
	//     description: user ID to get information
	//     required: true
	//     type: string
	//
	// security:
	// - apiKey: []
	// responses:
	//  200: User Returns a user for a given ID.
	//  404:
	// 	 description: User not found error if there is no user for the given ID
	//	500:
	//	 description: An internal server error occurred processing the request
	c.mux.HandleFunc("/user/{id}", c.GetUser).Methods(http.MethodGet)

	// swagger:route POST /user Users newUser
	// Creates a new User.
	//
	// Parameters:
	//   + name: userId
	//     in: path
	//     description: user ID to get information
	//     required: false
	//     type: string
	//   + name: userCreateRequest
	//     in: body
	//     description: User request object data
	//     required: true
	//     type: User
	//
	// security:
	// - apiKey: []
	// responses:
	//  200: User Returns created user object.
	//  400:
	//   description: not found error if there is no user for the given ID
	//	500:
	//   description: An internal server error occurred processing the request
	c.mux.HandleFunc("/user", c.NewUser).Methods(http.MethodPost)

	// swagger:route PUT /user/{userId} Users updateUser
	// Updates an existing User.
	//
	// Parameters:
	//   + name: userId
	//     in: path
	//     description: user ID to update
	//     required: true
	//     type: string
	//   + name: userUpdateRequest
	//     in: body
	//     description: User update request object data
	//     required: true
	//     type: User
	//
	// security:
	// - apiKey: []
	// responses:
	//  200: User
	//   description: Returns the updated user object
	//  400:
	//   description: Invalid user object request
	//	500:
	//   description: An internal server error occurred processing the request
	c.mux.HandleFunc("/user", c.UpdateUser).Methods(http.MethodPut, http.MethodPatch)

	// swagger:route DELETE /user/{id} Users deleteUser
	// Removes an existing User for a given ID.
	//
	// Parameters:
	//   + name: userId
	//     in: path
	//     description: ID value of the user to be removed
	//     required: false
	//     type: string
	//
	// security:
	// - apiKey: []
	// responses:
	//  200:
	//   description: User was successfully removed
	//  404:
	//   description: User not found error if there is no user for the given ID
	//	500:
	//   description: An internal server error occurred processing the request
	c.mux.HandleFunc("/user/{id}", c.DeleteUser).Methods(http.MethodDelete)

	c.mux.HandleFunc("/users", c.FindUsers).Methods(http.MethodGet)
}
