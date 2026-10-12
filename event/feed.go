package event

const FeedSchemaVersion = "lumi.event.feed.v1"

// Feed is the HTTP response envelope for a source that publishes event
// candidates. Each candidate retains its own source_ref for traceability.
type Feed struct {
	SchemaVersion string      `json:"schema_version"`
	Source        string      `json:"source"`
	GeneratedAt   string      `json:"generated_at,omitempty"`
	Events        []Candidate `json:"events"`
}
