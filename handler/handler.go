package handler

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if dest, ok := pathsToUrls[url]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

// Paths to store the paths to URLs
type Paths []struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYAML(yml []byte) (Paths, error) {
	paths := Paths{}
	err := yaml.Unmarshal(yml, &paths)

	// fmt.Printf("%v\n", paths)
	if err != nil {
		panic(err)
	}
	return paths, nil
}

func buildMap(parsedYaml Paths) map[string]string {
	pathMap := make(map[string]string, len(parsedYaml))
	for i := 0; i < len(parsedYaml); i++ {
		// fmt.Printf("%v %v -> %v\n", i+1, parsedYaml[i].Path, parsedYaml[i].URL)
		pathMap[parsedYaml[i].Path] = parsedYaml[i].URL
	}

	return pathMap
}
