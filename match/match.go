// Package match contains the versioned event matching decision contract.
package match

const DecisionSchemaVersion = "lumi.event.match_decision.v1"

type Status string

const (
	StatusMatched   Status = "matched"
	StatusAmbiguous Status = "ambiguous"
	StatusUnmatched Status = "unmatched"
)

type FeatureScores struct {
	Title      float64 `json:"title"`
	Date       float64 `json:"date"`
	Venue      float64 `json:"venue"`
	Performers float64 `json:"performers"`
}

type CandidateScore struct {
	CandidateID string        `json:"candidate_id"`
	Score       float64       `json:"score"`
	Features    FeatureScores `json:"features"`
}

type Alert struct {
	Type       string   `json:"type"`
	Reason     string   `json:"reason"`
	Candidates []string `json:"candidates,omitempty"`
}

type Decision struct {
	SchemaVersion       string           `json:"schema_version"`
	QueryID             string           `json:"query_id"`
	Status              Status           `json:"status"`
	SelectedCandidateID string           `json:"selected_candidate_id,omitempty"`
	TopScore            float64          `json:"top_score"`
	SecondScore         float64          `json:"second_score"`
	ScoreMargin         float64          `json:"score_margin"`
	Candidates          []CandidateScore `json:"candidates"`
	Alert               *Alert           `json:"alert,omitempty"`
}
