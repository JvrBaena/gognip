package types

/*
Rule ...
*/
type Rule struct {
	Value string `json:"value"`
	Tag   string `json:"tag,omitempty"`
	ID    string `json:"id,omitempty"`
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
	Summary ruleResponseSummary  `json:"summary"`
	Detail  []ruleResponseDetail `json:"detail"`
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
	addRuleResponseDetail
	removeRuleResponseDetail
}

type addRuleResponseDetail struct {
	Rule    Rule   `json:"rule,omitempty"`
	Created bool   `json:"created,omitempty"`
	Message string `json:"message,omitempty"`
}

type removeRuleResponseDetail struct {
	Rule    Rule   `json:"rule,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
	Message string `json:"message,omitempty"`
}