// Package event contains the versioned cross-system event contract.
package event

const CandidateSchemaVersion = "lumi.event.candidate.v1"

// SourceRef preserves the source identity needed for re-fetching and review.
type SourceRef struct {
	Source     string `json:"source"`
	EventID    string `json:"event_id"`
	URL        string `json:"url,omitempty"`
	CalendarID string `json:"calendar_id,omitempty"`
}

// Candidate is one source-normalized event before canonical merging.
// LocalDate is the date in the event's declared timezone and is the stable
// calendar feature used by the initial matcher PoC.
type Candidate struct {
	SchemaVersion string            `json:"schema_version"`
	ID            string            `json:"id"`
	SourceRef     SourceRef         `json:"source_ref"`
	Title         string            `json:"title"`
	LocalDate     string            `json:"local_date"`
	StartsAt      string            `json:"starts_at,omitempty"`
	EndsAt        string            `json:"ends_at,omitempty"`
	Timezone      string            `json:"timezone,omitempty"`
	Venue         string            `json:"venue,omitempty"`
	Performers    []string          `json:"performers,omitempty"`
	Extensions    map[string]string `json:"extensions,omitempty"`
}
