package users

import (
	"net/http"

	"github.com/andyglass/go-restapi-template/app/logger"
	"github.com/andyglass/go-restapi-template/app/responses"
	"github.com/go-chi/render"
)

var log = logger.Logger

func listUsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, responses.JSONResponse{
			Data: "pong",
		})
	}
}

func getUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, responses.JSONResponse{
			Data: "pong",
		})
	}
}

func createUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, responses.JSONResponse{
			Data: "pong",
		})
	}
}

func deleteUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, responses.JSONResponse{
			Data: "pong",
		})
	}
}
