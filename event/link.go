package event

// Link is a URL associated with an event candidate.
//
// SourceRef.URL remains the canonical URL of the source record. Links are
// additional URLs discovered while normalizing that record, such as a ticket
// page embedded in an event description.
type Link struct {
	Kind   string `json:"kind"`
	URL    string `json:"url"`
	Label  string `json:"label,omitempty"`
	Origin string `json:"origin,omitempty"`
}

const (
	LinkKindTicket   = "ticket"
	LinkKindOfficial = "official"
	LinkKindRelated  = "related"
	LinkKindSocial   = "social"
)
