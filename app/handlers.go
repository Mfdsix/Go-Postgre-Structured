package app

import (
	"fmt"
	"go-postgre/app/models"
	"log"
	"net/http"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to GoPostgre")
	}
}

func (a *App) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.RPost{}
		err := parse(w, r, &req)

		if err != nil {
			log.Printf("Cannot parse post body, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		p := &models.DPost{
			ID:      0,
			Title:   req.Title,
			Content: req.Content,
			Author:  req.Author,
		}

		// save to DB
		err = a.DB.CreatePost(p)
		if err != nil {
			log.Printf("Cannot save post in DB, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := models.MapPostToJSON(p)
		sendResponse(w, r, resp, http.StatusOK)
	}
}
