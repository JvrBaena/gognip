package types

/*
Rule ...
*/
type Rule struct {
	Value string `json:"value"`
	Tag   string `json:"tag,omitempty"`
	ID    int64  `json:"id,omitempty"`
}

/*
RuleRequest ...
*/
type RuleRequest struct {
	Rules []*Rule `json:"rules"`
}

/*
RuleRequestResponse ...
*/
type RuleRequestResponse struct {
	Summary ruleResponseSummary  `json:"summary,omitempty"`
	Rules   []Rule               `json:"rules,omitempty"`
	Detail  []ruleResponseDetail `json:"detail,omitempty"`
	Sent    string               `json:"sent,omitempty"`
}

type ruleResponseSummary struct {
	addRuleSummary
	removeRuleSummary
}

type addRuleSummary struct {
	Created    int `json:"created,omitempty"`
	NotCreated int `json:"not_created,omitempty"`
}

type removeRuleSummary struct {
	Deleted    int `json:"deleted,omitempty"`
	NotDeleted int `json:"not_deleted,omitempty"`
}

type ruleResponseDetail struct {
	Rule    Rule   `json:"rule,omitempty"`
	Message string `json:"message,omitempty"`
	Created bool   `json:"created,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}
