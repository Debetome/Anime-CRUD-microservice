package handlers

import (
	"microservice/models"
	"net/http"
	"fmt"
	"context"
)

type KeyAnime struct{}

func (self Animes) MiddlewareValidateAnime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		anime := models.Anime{}

		err := anime.FromJSON(r.Body)
		if err != nil {
			self.logger.Println("[ERROR] deserializing anime", err)
			http.Error(rw, "Error reading anime", http.StatusBadRequest)
			return
		}

		err = anime.Validate()
		if err != nil {
			self.logger.Println("[ERROR] validating anime", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating anime: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyAnime{}, anime)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}