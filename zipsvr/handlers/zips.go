package handlers

import "github.com/johnlawsharrison/info344-in-class/zipsvr/models"
import "net/http"
import "strings"
import "encoding/json"

type CityHandler struct {
	PathPrefix string
	Index      models.ZipIndex
}

// receiver function
func (ch *CityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// URL: /zips/[city-name]
	cityName := strings.ToLower(r.URL.Path[len(ch.PathPrefix):])
	if len(cityName) == 0 {
		http.Error(w, "please provide a city name", http.StatusBadRequest)
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add(headerContentType, contentTypeJSON)
	zips := ch.Index[cityName]
	json.NewEncoder(w).Encode(zips)
}
