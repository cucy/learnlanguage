package handlers

import "net/http"

func FoodHandler(w http.ResponseWriter, r http.Request) {
	fooId := r.Context().Value("fooId").(string)
	w.Write([]byte(fooId))
}
