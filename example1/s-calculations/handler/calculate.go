package handler

import (
	"context"
	"io"
	"time"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
	"github.com/auvn/go-examples/example1/frwk-core/encoding"
	"github.com/auvn/go-examples/example1/frwk-core/transport/event/eventutil"
	"github.com/auvn/go-examples/example1/s-calculations/calculationsevent"
	"github.com/shopspring/decimal"
)

type Handlers struct {
	Events *eventutil.Publisher
}

func (h *Handlers) CalculateTrip(ctx context.Context, r io.Reader) error {
	var req struct {
		TripID id.ID

		Distance int
		Duration time.Duration
	}

	if err := encoding.UnmarshalReader(r, &req); err != nil {
		return err
	}

	dist := decimal.NewFromFloat(float64(req.Distance))
	dur := decimal.NewFromFloat(req.Duration.Minutes())
	rate := decimal.NewFromFloat(float64(time.Now().Minute()))

	return h.Events.PublishEvent(ctx,
		calculationsevent.TypeBreakdownCalculated,
		calculationsevent.BreakdownCalculated{
			TripID: req.TripID,
			Total:  dist.Add(dur).Mul(rate).Round(2),
		})
}
