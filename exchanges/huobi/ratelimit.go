package huobi

import (
	"context"
	"time"

	"github.com/otter-trade/coin-exchange-api/exchanges/request"
	"golang.org/x/time/rate"
)

const (
	// Huobi rate limits per API Key
	huobiSpotRateInterval = time.Second * 1
	huobiSpotRequestRate  = 7

	huobiFuturesRateInterval    = time.Second * 3
	huobiFuturesAuthRequestRate = 30
	// Non market-request public interface rate
	huobiFuturesUnAuthRequestRate    = 60
	huobiFuturesTransferRateInterval = time.Second * 3
	huobiFuturesTransferReqRate      = 10

	huobiSwapRateInterval      = time.Second * 3
	huobiSwapAuthRequestRate   = 30
	huobiSwapUnauthRequestRate = 60

	huobiFuturesAuth request.EndpointLimit = iota
	huobiFuturesUnAuth
	huobiFuturesTransfer
	huobiSwapAuth
	huobiSwapUnauth
)

// RateLimit implements the request.Limiter interface
type RateLimit struct {
	Spot          *rate.Limiter
	FuturesAuth   *rate.Limiter
	FuturesUnauth *rate.Limiter
	SwapAuth      *rate.Limiter
	SwapUnauth    *rate.Limiter
	FuturesXfer   *rate.Limiter
}

// Limit limits outbound requests
func (r *RateLimit) Limit(ctx context.Context, f request.EndpointLimit) error {
	switch f {
	// TODO: Add futures and swap functionality
	case huobiFuturesAuth:
		return r.FuturesAuth.Wait(ctx)
	case huobiFuturesUnAuth:
		return r.FuturesUnauth.Wait(ctx)
	case huobiFuturesTransfer:
		return r.FuturesXfer.Wait(ctx)
	case huobiSwapAuth:
		return r.SwapAuth.Wait(ctx)
	case huobiSwapUnauth:
		return r.SwapUnauth.Wait(ctx)
	default:
		// Spot calls
		return r.Spot.Wait(ctx)
	}
}

// SetRateLimit returns the rate limit for the exchange
func SetRateLimit() *RateLimit {
	return &RateLimit{
		Spot:          request.NewRateLimit(huobiSpotRateInterval, huobiSpotRequestRate),
		FuturesAuth:   request.NewRateLimit(huobiFuturesRateInterval, huobiFuturesAuthRequestRate),
		FuturesUnauth: request.NewRateLimit(huobiFuturesRateInterval, huobiFuturesUnAuthRequestRate),
		SwapAuth:      request.NewRateLimit(huobiSwapRateInterval, huobiSwapAuthRequestRate),
		SwapUnauth:    request.NewRateLimit(huobiSwapRateInterval, huobiSwapUnauthRequestRate),
		FuturesXfer:   request.NewRateLimit(huobiFuturesTransferRateInterval, huobiFuturesTransferReqRate),
	}
}
