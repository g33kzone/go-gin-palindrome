package model

// NoteRequest - struct symbolises the Note request object
type NoteRequest struct {
	Message string `json:"message"`
}

// NoteResponse - struct symbolises the Note response object
type NoteResponse struct {
	ID           uint   `json:"id"`
	Message      string `json:"message"`
	IsPalindrome bool   `json:"ispalindrome"`
}

// NotesResponse - struct symbolises the list of NotesResponse
type NotesResponse struct {
	Notes []NoteResponse
}

// ErrorResponse - struct for error response
type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}
