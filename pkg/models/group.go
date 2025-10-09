package models

import "time"

type Group struct {
	Schemas     []string `json:"schemas"` //urn:ietf:params:scim:schemas:core:2.0:Group
	ID          string   `json:"id"`
	DisplayName string   `json:"displayName"`
	Members     []struct {
		Value   string `json:"value,omitempty"`
		Ref     string `json:"$ref,omitempty"`
		Display string `json:"display,omitempty"`
	} `json:"members,omitempty"`
	Meta struct {
		ResourceType string    `json:"resourceType"`
		Created      time.Time `json:"created"`
		LastModified time.Time `json:"lastModified"`
		Version      string    `json:"version"`
		Location     string    `json:"location"`
	} `json:"meta"`
}
