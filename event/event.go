// Package event contains the versioned cross-system event contract.
package event

const CandidateSchemaVersion = "lumi.event.candidate.v1"

// SourceRef preserves the source identity needed for re-fetching and review.
type SourceRef struct {
	Source       string `json:"source"`
	RecordID     string `json:"record_id"`
	URL          string `json:"url,omitempty"`
	CollectionID string `json:"collection_id,omitempty"`
}

// EntityRef points to a canonical or source-local person, group, venue, or
// organizer without forcing every input system to resolve identities first.
type EntityRef struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

// Candidate is one source-normalized event before canonical merging.
// LocalDate is the date in the event's declared timezone and is the stable
// calendar feature used by the initial matcher PoC.
type Candidate struct {
	SchemaVersion string         `json:"schema_version"`
	ID            string         `json:"id"`
	SourceRef     SourceRef      `json:"source_ref"`
	Title         string         `json:"title"`
	LocalDate     string         `json:"local_date"`
	StartsAt      string         `json:"starts_at,omitempty"`
	EndsAt        string         `json:"ends_at,omitempty"`
	Timezone      string         `json:"timezone,omitempty"`
	Venue         string         `json:"venue,omitempty"`
	Performers    []string       `json:"performers,omitempty"`
	PerformerRefs []EntityRef    `json:"performer_refs,omitempty"`
	OrganizerRefs []EntityRef    `json:"organizer_refs,omitempty"`
	Links         []Link         `json:"links,omitempty"`
	Extensions    map[string]any `json:"extensions,omitempty"`
}
