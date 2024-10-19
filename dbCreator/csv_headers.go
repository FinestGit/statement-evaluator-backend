package dbcreator

import (
	"net/http"
	"unicode"

	"github.com/finestgit/statement-evaluator-backend/models"
	"github.com/gin-gonic/gin"
)

func RegisterHeaderRoutes(router *gin.Engine) {
	router.POST("/dbCreator/csvHeaders/transform", transformCSVheaders)
}

func transformCSVheaders(context *gin.Context) {
	var csvHeaders models.CSVHeaders
	err := context.ShouldBindJSON(&csvHeaders)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body."})
		return
	}

	var modifiedHeaders []string

	for _, header := range csvHeaders.CSVHeaders {
		var modifiedRunes []rune
		var charBefore rune

		for index, char := range header {
			if unicode.IsSpace(char) && index != 0 {
				modifiedRunes = append(modifiedRunes, '_')
			} else if unicode.IsUpper(char) {
				if index == 0 {
					modifiedRunes = append(modifiedRunes, unicode.ToLower(char))
				} else if unicode.IsSpace(charBefore) {
					modifiedRunes = append(modifiedRunes, unicode.ToLower(char))
				} else {
					modifiedRunes = append(modifiedRunes, '_')
					modifiedRunes = append(modifiedRunes, unicode.ToLower(char))
				}
			} else {
				modifiedRunes = append(modifiedRunes, char)
			}
			charBefore = char
		}

		modifiedHeaders = append(modifiedHeaders, string(modifiedRunes))
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully transformed", "modifiedHeaders": modifiedHeaders})
}
