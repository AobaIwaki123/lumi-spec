// Package entity contains the shared idol and group profile contract.
package entity

import "github.com/AobaIwaki123/lumi-spec/event"

const SchemaVersion = "lumi.entity.v1"

type Kind string

const (
	KindIdol      Kind = "idol"
	KindGroup     Kind = "group"
	KindVenue     Kind = "venue"
	KindOrganizer Kind = "organizer"
)

type AssetKind string

const (
	AssetIcon  AssetKind = "icon"
	AssetLogo  AssetKind = "logo"
	AssetImage AssetKind = "image"
)

type Asset struct {
	Kind     AssetKind `json:"kind"`
	URL      string    `json:"url"`
	Alt      string    `json:"alt,omitempty"`
	Width    int       `json:"width,omitempty"`
	Height   int       `json:"height,omitempty"`
	MimeType string    `json:"mime_type,omitempty"`
}

type SocialLink struct {
	Platform string `json:"platform"`
	URL      string `json:"url"`
	Handle   string `json:"handle,omitempty"`
	Label    string `json:"label,omitempty"`
}

// Record is a cross-system profile record for an idol, group, venue, or
// organizer. Source refs are retained so name resolution can be revisited.
type Record struct {
	SchemaVersion string            `json:"schema_version"`
	ID            string            `json:"id"`
	Kind          Kind              `json:"kind"`
	Name          string            `json:"name"`
	Aliases       []string          `json:"aliases,omitempty"`
	Assets        []Asset           `json:"assets,omitempty"`
	SocialLinks   []SocialLink      `json:"social_links,omitempty"`
	SourceRefs    []event.SourceRef `json:"source_refs"`
	Extensions    map[string]any    `json:"extensions,omitempty"`
}
