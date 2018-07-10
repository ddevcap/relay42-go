package relay42

const (
	RuleTypeAnd = "AND"
	RuleTypeOr  = "OR"
)

type InteractionQuery struct {
}

type InteractionRule struct {
	ruleType string
	Key      string
	Value    string
	Operator string
}

func (i *InteractionQuery) And(andRules ...*InteractionRule) *InteractionQuery {
	for _, rule := range andRules {
		rule.ruleType = RuleTypeAnd
	}
	return i
}

func (i *InteractionQuery) Or(orRules ...*InteractionRule) *InteractionQuery {
	for _, rule := range orRules {
		rule.ruleType = RuleTypeOr
	}
	return i
}

func NewInteractionQuery() *InteractionQuery {
	return &InteractionQuery{}
}
