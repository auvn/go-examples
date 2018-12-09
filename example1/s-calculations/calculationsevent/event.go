package calculationsevent

import (
	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
	"github.com/shopspring/decimal"
)

const (
	TypeBreakdownCalculated = "calculations_breakdown_calculated"
)

type BreakdownCalculated struct {
	TripID id.ID

	Total decimal.Decimal
}
