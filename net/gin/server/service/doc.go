package service

// swagger:parameters UpdateUserResponseWrapper
type UpdateUserRequest struct {
	// in: body
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// Update User Info
//
// swagger:response UpdateUserResponseWrapper
type UpdateUserResponseWrapper struct {
	// in: body
	Score string `json:"score"`
}

// swagger:route POST /users users UpdateUserResponseWrapper
//
// Update User
//
// This will update user info
//
//     Responses:
//       200: UpdateUserResponseWrapper
