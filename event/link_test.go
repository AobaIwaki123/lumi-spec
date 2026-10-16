package event

import (
	"encoding/json"
	"testing"
)

func TestCandidateLinksMarshal(t *testing.T) {
	candidate := Candidate{
		SchemaVersion: CandidateSchemaVersion,
		ID:            "example:1",
		SourceRef: SourceRef{
			Source:   "example",
			RecordID: "1",
			URL:      "https://example.test/events/1",
		},
		Title:     "Example",
		LocalDate: "2026-10-16",
		Links: []Link{{
			Kind:   LinkKindTicket,
			URL:    "https://example.test/tickets/1",
			Label:  "一般販売",
			Origin: "description",
		}},
	}

	data, err := json.Marshal(candidate)
	if err != nil {
		t.Fatalf("marshal candidate: %v", err)
	}
	var decoded Candidate
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal candidate: %v", err)
	}
	if len(decoded.Links) != 1 || decoded.Links[0].Kind != LinkKindTicket {
		t.Fatalf("links = %#v", decoded.Links)
	}
}
