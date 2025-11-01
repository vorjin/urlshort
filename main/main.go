package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// opening yaml file as a flag
	yamlFilePath := flag.String("yaml", "paths.yaml", "path to the yaml file to open")
	flag.Parse()
	yamlFile, err := os.ReadFile(*yamlFilePath)
	if err != nil {
		fmt.Printf("Error opening the file: %s", *yamlFilePath)
		return
	}

	/* yaml := `
	- path: /urlshort
	  url: https://github.com/gophercises/urlshort
	- path: /urlshort-final
	  url: https://github.com/gophercises/urlshort/tree/solution
	` */

	yamlHandler, err := urlshort.YAMLHandler([]byte(yamlFile), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	err = http.ListenAndServe(":8080", yamlHandler)
	if err != nil {
		log.Fatal("Server failed: ", err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
