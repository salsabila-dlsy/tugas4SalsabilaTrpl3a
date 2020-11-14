package middleware

import (
	"time"

	"context"

	// "pnp/Framework/git/order/services"
	cm "tugas4SalsabilaTrpl3a/Framework/git/order/common"

	log "github.com/Sirupsen/logrus"
)

//BasicMiddleware is
func BasicMiddleware() services.ServiceMiddleware {
	return func(next services.PaymentServices) services.PaymentServices {
		return BasicMiddlewareStruct{next}
	}
}

//BasicMiddlewareStruct is
type BasicMiddlewareStruct struct {
	services.PaymentServices
}

func (mw BasicMiddlewareStruct) OrderHandler(ctx context.Context, request cm.Message) cm.Message {

	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("OrderHandler ends")
	}(time.Now())

	log.WithField("request", request).Info("OrderHandler begins")

	return mw.PaymentServices.OrderHandler(ctx, request)

}
