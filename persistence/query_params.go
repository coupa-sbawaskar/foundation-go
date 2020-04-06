package persistence

type QueryParams struct {
	Operands []QueryExpression
	Limit    uint64
	Offset   uint64
	Order    []OrderStatement
}

type QueryOperatorType string

const (
	QUERY_OPERATOR_EQ          = "EQ"
	QUERY_OPERATOR_NEQ         = "NEQ"
	QUERY_OPERATOR_CONTAINS    = "CONTAINS"
	QUERY_OPERATOR_IN          = "IN"
	QUERY_OPERATOR_GT          = "GT"
	QUERY_OPERATOR_GTE         = "GTE"
	QUERY_OPERATOR_LT          = "LT"
	QUERY_OPERATOR_LTE         = "LTE"
	QUERY_OPERATOR_STARTS_WITH = "STARTS_WITH"
	QUERY_OPERATOR_ENDS_WITH   = "ENDS_WITH"
)

type QueryExpression struct {
	Key      string
	Operator QueryOperatorType
	Value    string
}

type OrderStatement struct {
	ColumnName string
	Direction  OrderDirection
}

type OrderDirection string

const (
	ORDER_DIRECTION_ASC  = "asc"
	ORDER_DIRECTION_DESC = "desc"
)
