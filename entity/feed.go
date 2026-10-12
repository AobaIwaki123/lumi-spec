package entity

const FeedSchemaVersion = "lumi.entity.feed.v1"

// Feed is the HTTP response envelope for a source that publishes entities.
type Feed struct {
	SchemaVersion string   `json:"schema_version"`
	Source        string   `json:"source"`
	GeneratedAt   string   `json:"generated_at,omitempty"`
	Entities      []Record `json:"entities"`
}
