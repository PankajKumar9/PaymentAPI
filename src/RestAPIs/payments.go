package RestAPIs

import (
	"github.com/PankajKumar9/PaymentAPI/src/conf"
	routing "github.com/qiangxue/fasthttp-routing"
)

func InitApis(r *routing.Router) {
	//TODO: paths to be in cosntants package

	r.Post("/api/signup", func(c *routing.Context) error {
		return conf.Signup(string(c.PostBody()), c)
		//return somefolder.somefunc(string(c.PostBody()), c)
	})
	r.Post("/api/credit", func(c *routing.Context) error {
		return conf.Credit(string(c.PostBody()), c)
		//return somefolder.somefunc(string(c.PostBody()), c)
	})
	r.Post("/api/debit", func(c *routing.Context) error {
		return conf.Debit(string(c.PostBody()), c)
		//return somefolder.somefunc(string(c.PostBody()), c)
	})
	r.Post("/api/send", func(c *routing.Context) error {
		return conf.Send(string(c.PostBody()), c)
		//return somefolder.somefunc(string(c.PostBody()), c)
	})
	r.Post("/api/refund", func(c *routing.Context) error {
		return conf.RefundResponse(string(c.PostBody()), c)
		//return somefolder.somefunc(string(c.PostBody()), c)
	})
	r.Get("/api/history", func(c *routing.Context) error {
		conf.History(c)
		return nil
	})
}
