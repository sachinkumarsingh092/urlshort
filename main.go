package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sachinkumarsingh092/urlshort/handler"
)

const (
	yamlFlag            = "yaml"
	yamlFlagDefVal      = "urls.yaml"
	yamlFlagDescription = "Sets the YAML file to create URL map."
)

// Flagger is an interface to configure various flags.
type Flagger interface {
	StringVar(p *string, name, defval, description string)
}

type urlshortFlagger struct{}

func (u *urlshortFlagger) StringVar(p *string, name, defval, description string) {
	flag.StringVar(p, name, defval, description)
}

var yaml string

// ConfigFlag configures the flags.
func ConfigFlag(flagger Flagger) {
	flagger.StringVar(&yaml, yamlFlag, yamlFlagDefVal, yamlFlagDescription)
}

var pathsToUrls = map[string]string{
	"/http-godoc": "https://godoc.org/net/http",
	"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
}

func createMapHandler() http.HandlerFunc {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	return mapHandler
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func createYamlHnadler(yaml string, fallback http.HandlerFunc) http.HandlerFunc {
	// Read the entire file `yaml` and save its content to yamlFile
	yamlFile, err := ioutil.ReadFile(yaml)

	if err != nil {
		panic(err)
	}

	// Build the YAMLHandler using the mapHandler as the fallback
	yamlHandler, err := handler.YAMLHandler(yamlFile, fallback)
	if err != nil {
		panic(err)
	}
	return yamlHandler
}

func main() {
	flagger := &urlshortFlagger{}
	ConfigFlag(flagger)
	flag.Parse()

	mapHandler := createMapHandler()
	yamlHandler := createYamlHnadler(yaml, mapHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}
