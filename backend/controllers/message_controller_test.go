package controllers

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/techotron/online-quote-book/backend/constants"
)

func TestMessageHandler(t *testing.T) {
	actualResponse := MessageHandler(constants.ErrorNoRowsFound)
	assert.Equal(t, gin.H{"message": constants.ErrorNoRowsFound}, actualResponse)
}
