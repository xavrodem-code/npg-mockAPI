package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
)
type App struct {
	cards []card
}

type card struct {
	ID	string `json:id`
	Strength int `json:strength`
	Speed int `json:speed`
	Skills int `json:skills`
	Stealth int `json:stealth`
	Stamina int `json:stamina`
	Sexiness int `json:stamina`
}

func (app *App) getCard(context *gin.Context) {
	id := context.Param("id")
	card, err := app.getCardById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"Card not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, card)
}

func (app *App) getCardById(id string) (*card, error) {
	for i, t := range app.cards {
		if t.ID == id {
			return &app.cards[i], nil
		}
	}
	return nil, errors.New("card not found")
}

func main() {
	app := &App {
		cards: []card{
			{ID:"1", Strength: 34, Speed: 23, Skills: 30, Stealth: 32, Stamina: 36, Sexiness: 32},
			{ID:"2", Strength: 31, Speed: 23, Skills: 29, Stealth: 27, Stamina: 22, Sexiness: 10},
		},
	}

	r := gin.Default()

	r.GET("/cards/:id", app.getCard)

	r.Run("localhost:9090")
}