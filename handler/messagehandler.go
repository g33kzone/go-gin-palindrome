package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/g33kzone/go-message-palindrome/db"
	"github.com/g33kzone/go-message-palindrome/model"
	"github.com/g33kzone/go-message-palindrome/service"
	"github.com/gin-gonic/gin"
)

// SaveMessageHandler - Handler to save messages received
// Authenticate godoc
// @Summary Save message to postgres db
// @Description Captures messages received and saves to the postgresql database
// @ID note-post-g33kzone
// @Accept json
// @Produce json
// @Param message body model.NoteRequest true "Message"
// @Success 200 {object} model.Note
// @Failure 401 {object} model.ErrorResponse
// @Router /notes [post]
func SaveMessageHandler(c *gin.Context, db *db.Db) {
	var noteRequest model.NoteRequest

	if err := c.ShouldBindJSON(&noteRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chkFlag := service.IsPalindrome(strings.ToLower(noteRequest.Message))
	note := model.Note{Message: noteRequest.Message, IsPalindrome: chkFlag}
	id, err := db.SaveMessage(note)
	if err != nil {
		panic(err)
	}
	note = model.Note{ID: id, Message: noteRequest.Message, IsPalindrome: chkFlag}

	c.JSON(http.StatusOK, note)
}

// FetchMessageListHandler - Handler to list all messages received
// Authenticate godoc
// @Summary Fetch all recent messages saved to database
// @Description Captures messages received and saves to the postgresql database
// @ID note-list-get-g33kzone
// @Accept json
// @Produce json
// @Success 200 {object} model.NotesResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /notes [get]
func FetchMessageListHandler(c *gin.Context, db *db.Db) {

	notesResponse, err := db.FetchAllMessages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{ErrorMessage: err.Error()})
	}

	c.JSON(http.StatusOK, notesResponse)
}

// FetchMessageHandler - Handler to fetch a single message based on id
// Authenticate godoc
// @Summary Fetch a single message based on id
// @Description Captures messages received and saves to the postgresql database
// @ID note-get-g33kzone
// @Accept json
// @Produce json
// @Param id path uint true "ID"
// @Success 200 {object} model.NoteResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /notes/{id} [get]
func FetchMessageHandler(c *gin.Context, db *db.Db) {

	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{ErrorMessage: err.Error()})
	}
	notesResponse, err := db.FetchMessage(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{ErrorMessage: err.Error()})
	}

	c.JSON(http.StatusOK, notesResponse)
}

// DeleteMessageHandler - Handler to delete a single message based on id
// Authenticate godoc
// @Summary Delete a single message based on id
// @Description Delete messages from postgres db based on id received from Delete API
// @ID note-delete-g33kzone
// @Accept json
// @Produce json
// @Param id path uint true "ID"
// @Success 200 {object} string
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /notes/{id} [delete]
func DeleteMessageHandler(c *gin.Context, db *db.Db) {

	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{ErrorMessage: err.Error()})
	}

	count, err := db.DeleteMessage(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{ErrorMessage: err.Error()})
	}

	if count > 0 {
		c.JSON(http.StatusOK, fmt.Sprintf("Message with ID - %d deleted successfully", id))
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf("Oops...Message with ID - %d does not exists", id))
	}

}
