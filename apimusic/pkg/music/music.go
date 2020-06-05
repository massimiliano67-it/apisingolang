package music

import (
	"context"
	"time"
)

// Music defines the properties of a music to be listed
type Music struct {
	ID        string     `json:"ID"`
	Name      string     `json:"name,omitempty"`
	Artist    string     `json:"artist,omitempty"`
	Image     string     `json:"image,omitempty"`
	Year      string     `json:"year,omitempty"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}

// MusicRepository Repository provides access to the music storage
type MusicRepository interface {
	// CreateMusic saves a given music
	CreateMusic(ctx context.Context, music *Music) error
	// FetchMusics return all music saved in storage
	FetchMusics(ctx context.Context) ([]*Music, error)
	// DeleteMusic remove music with given ID
	DeleteMusic(ctx context.Context, ID string) error
	// UpdateMusic modify music with given ID and given new data
	UpdateMusic(ctx context.Context, ID string, music *Music) error
	// FetchMusicByID returns the music with given ID
	FetchMusicByID(ctx context.Context, ID string) (*Music, error)
}
