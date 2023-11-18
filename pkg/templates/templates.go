// load templates from templates.json file
package templates

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed templates.json
var templates []byte

type Template struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type Group struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Templates   []Template `json:"templates"`
}

type AllGroups struct {
	Groups []Group `json:"groups"`
}

// GetGroups returns the templates
func GetGroups() (AllGroups, error) {
	// decode json file
	var allTemplates AllGroups
	err := json.Unmarshal(templates, &allTemplates)
	if err != nil {
		log.Fatalf("unable to decode templates.json: %v", err)
		return AllGroups{}, err
	}
	return allTemplates, nil

}
