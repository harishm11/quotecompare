package ruletables

type DiscountRules struct {
	StateCode        string
	ProductCode      string
	TransactionCode  string
	DiscountCode     string
	RuleId           string
	RuleSequence     int
	RuleGroupCode    string
	TestCriteriaCode string
	TestOperator     string
	TestValue        string
	ActionCode       string
	MessageCode      string
	Scope            string
}
