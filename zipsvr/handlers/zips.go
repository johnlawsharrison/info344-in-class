package handlers

import "github.com/johnlawsharrison/info344-in-class/zipsvr/models"

type CityHandler struct {
	PathPrefix string
	Index      models.ZipIndex
}

// receiver function
func (ch *CityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// URL: /zips/[city-name]

}
