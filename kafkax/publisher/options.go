package publisher

import "github.com/txix-open/isp-kit/kafkax"

type Option func(p *Publisher)

func WithMiddlewares(mws ...Middleware) Option {
	return func(p *Publisher) {
		p.Middlewares = append(p.Middlewares, mws...)
	}
}

func WithObserver(observer kafkax.Observer) Option {
	return func(c *Publisher) {
		c.observer = observer
	}
}
