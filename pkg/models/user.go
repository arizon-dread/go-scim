package models

import "time"

type User struct {
	Schemas    []string `json:"schemas"` //urn:ietf:params:scim:schemas:core:2.0:User
	ID         string   `json:"id"`
	ExternalID string   `json:"externalId,omitempty"`
	UserName   string   `json:"userName"`
	Name       struct {
		Formatted       string `json:"formatted,omitempty"`
		FamilyName      string `json:"familyName,omitempty"`
		GivenName       string `json:"givenName,omitempty"`
		MiddleName      string `json:"middleName,omitempty"`
		HonorificPrefix string `json:"honorificPrefix,omitempty"`
		HonorificSuffix string `json:"honorificSuffix,omitempty"`
	} `json:"name"`
	DisplayName string `json:"displayName,omitempty"`
	NickName    string `json:"nickName,omitempty"`
	ProfileURL  string `json:"profileUrl,omitempty"`
	Emails      []struct {
		Value   string `json:"value,omitempty"`
		Type    string `json:"type,omitempty"`
		Primary bool   `json:"primary,omitempty"`
	} `json:"emails,omitempty"`
	Addresses []struct {
		Type          string `json:"type,omitempty"`
		StreetAddress string `json:"streetAddress,omitempty"`
		Locality      string `json:"locality,omitempty"`
		Region        string `json:"region,omitempty"`
		PostalCode    string `json:"postalCode,omitempty"`
		Country       string `json:"country,omitempty"`
		Formatted     string `json:"formatted,omitempty"`
		Primary       bool   `json:"primary,omitempty"`
	} `json:"addresses,omitempty"`
	PhoneNumbers []struct {
		Value string `json:"value,omitempty"`
		Type  string `json:"type,omitempty"`
	} `json:"phoneNumbers,omitempty"`
	Ims []struct {
		Value string `json:"value,omitempty"`
		Type  string `json:"type,omitempty"`
	} `json:"ims,omitempty"`
	Photos []struct {
		Value string `json:"value,omitempty"`
		Type  string `json:"type,omitempty"`
	} `json:"photos,omitempty"`
	UserType          string `json:"userType,omitempty"`
	Title             string `json:"title,omitempty"`
	PreferredLanguage string `json:"preferredLanguage,omitempty"`
	Locale            string `json:"locale,omitempty"`
	Timezone          string `json:"timezone,omitempty"`
	Active            bool   `json:"active,omitempty"`
	Password          string `json:"password,omitempty"`
	Groups            []struct {
		Value   string `json:"value,omitempty"`
		Ref     string `json:"$ref,omitempty"`
		Display string `json:"display,omitempty"`
	} `json:"groups,omitempty"`
	X509Certificates []struct {
		Value string `json:"value,omitempty"`
	} `json:"x509Certificates,omitempty"`
	Meta struct {
		ResourceType string    `json:"resourceType"`
		Created      time.Time `json:"created"`
		LastModified time.Time `json:"lastModified"`
		Version      string    `json:"version"`
		Location     string    `json:"location"`
	} `json:"meta"`
}
