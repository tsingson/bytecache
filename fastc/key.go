package fastc

import (
	"bytes"

	"github.com/valyala/fasthttp"
)

// Key represents a unique identifier for a resource in the cache
type Key struct {
	method []byte
	uri    []byte
	args   *fasthttp.Args
	header fasthttp.RequestHeader
}

// NewKey returns a new Key instance
func NewKey(ctx *fasthttp.RequestCtx) Key {
	return Key{method: ctx.Method(), header: ctx.Request.Header, args: ctx.QueryArgs(), uri: ctx.RequestURI()}
}

func (k Key) String() []byte {
	b := &bytes.Buffer{}
	b.Write(k.uri)
	b.Write(k.method)

	if k.args.Len() > 0 {
		k.args.VisitAll(
			func(key, value []byte) {
				// log.Warn("args", zap.String("key", vtils.B2S(key)), zap.String("value", vtils.B2S(value)))

				b.Write([]byte("::"))
				b.Write(key)

				b.Write([]byte("::"))
				b.Write(value)
			})
	}

	k.header.VisitAll(func(key, value []byte) {
		// log.Info("requestHeader", zap.String("key", gotils.B2S(key)), zap.String("value", gotils.B2S(value)))
		// log.Warn("header", zap.String("key", vtils.B2S(key)), zap.String("value", vtils.B2S(value)))
		b.Write([]byte("::"))
		b.Write(key)

		b.Write([]byte("::"))
		b.Write(value)
	})

	return b.Bytes()
}
