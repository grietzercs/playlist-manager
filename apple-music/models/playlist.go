package applemmusic

import (
	"time"
)

// structs generated from GPT to save time with boilerplate
// Root is the top-level structure for unmarshaling the entire response.
type Root struct {
	Data []Playlist `json:"data"`
}

// Playlist represents each individual playlist.
type Playlist struct {
	ID         string        `json:"id"`
	Type       string        `json:"type"`
	Href       string        `json:"href"`
	Attributes PlaylistAttrs `json:"attributes"`
}

// PlaylistAttrs contains the detailed attributes of each playlist.
type PlaylistAttrs struct {
	PlayParams   PlayParams    `json:"playParams"`
	CanEdit      bool          `json:"canEdit"`
	Name         string        `json:"name"`
	Description  Description   `json:"description"`
	DateAdded    time.Time     `json:"dateAdded"`
	Artwork      Artwork       `json:"artwork"`
	IsPublic     bool          `json:"isPublic"`
	HasCatalog   bool          `json:"hasCatalog"`
}

// PlayParams holds information about the playlist's parameters.
type PlayParams struct {
	ID        string `json:"id"`
	Kind      string `json:"kind"`
	IsLibrary bool   `json:"isLibrary"`
	GlobalID  string `json:"globalId"`
}

// Description represents different language versions of a description.
type Description struct {
	Standard string `json:"standard"`
}

// Artwork holds the image metadata for the playlist.
type Artwork struct {
	Width  *int   `json:"width"`  // Use pointer to handle null values
	Height *int   `json:"height"` // Use pointer to handle null values
	URL    string `json:"url"`
}