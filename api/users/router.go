package users

import (
	"github.com/go-chi/chi"
)

// GetRouter for users
func GetRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", listUsersHandler())
	r.Get("/", getUserHandler())
	r.Post("/", createUserHandler())
	r.Delete("/", deleteUserHandler())

	// r.Route("/", func(api chi.Router) {
	// 	api.Get("/*", s.whoamiHandler())
	// 	api.Put("/*", s.whoamiHandler())
	// 	api.Post("/*", s.whoamiHandler())
	// 	api.Patch("/*", s.whoamiHandler())
	// 	api.Delete("/*", s.whoamiHandler())
	// 	api.Options("/*", s.whoamiHandler())
	// })

	return r
}
