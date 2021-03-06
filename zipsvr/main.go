package main

import "fmt"
import "net/http"
import "log"
import "runtime"
import "encoding/json"
import "os"
import "strings"
import "github.com/johnlawsharrison/info344-in-class/zipsvr/models"
import "github.com/johnlawsharrison/info344-in-class/zipsvr/handlers"

const zipsPath = "/zips/"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Add("Content-Type", "text/plain")

	fmt.Fprintf(w, "Hello %s!", name)
}

func memoryHandler(w http.ResponseWriter, r *http.Request) {
	runtime.GC()
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func main() {

	addr := os.Getenv("ADDR")

	if len(addr) == 0 {
		addr = ":80"
	}
	zips, err := models.LoadZips("zips.csv")
	if err != nil {
		log.Fatalf("error loading zips: %v", err)
	}
	log.Printf("loaded %d zips", len(zips))

	// create a map of city names to zipcodes for that city
	cityIndex := models.ZipIndex{}
	for _, z := range zips {
		cityLower := strings.ToLower(z.City)
		cityIndex[cityLower] = append(cityIndex[cityLower], z)
	}

	//fmt.Println("Hello World!")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/memory", memoryHandler)
	fmt.Printf("server is listening at http://%s\n", addr)

	cityHandler := &handlers.CityHandler{
		Index:      cityIndex,
		PathPrefix: zipsPath,
	}
	mux.Handle(zipsPath, cityHandler)

	log.Fatal(http.ListenAndServe(addr, mux))
}
