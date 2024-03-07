AxMiddleware Processor
======================

```go

type Request struct {
    Value int
}

type Response struct {
    Result int
}

func main() {
    proc := NewProcessor[*Request, *Response]().WithLogger(log.Logger).WithMiddlewares(SumMiddleware, LogMiddleware).WithHandlers(func(ctx *Context[*Request, *Response]) {
        log.Debug().Msgf("result: %d", ctx.Response().Result)
    }).Build()

    var response Response
    proc.Process(nil, &Request{Value: 10}, &response)
}

```