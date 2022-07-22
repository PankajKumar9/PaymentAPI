package RestAPIs

import (
	"github.com/PankajKumar9/PaymentAPI/src/conf"
	"github.com/PankajKumar9/PaymentAPI/src/conf/constants"
	routing "github.com/qiangxue/fasthttp-routing"
)

func InitApis(r *routing.Router) {
	//TODO: paths to be in cosntants package

	r.Post(constants.DS+constants.API+constants.DS+constants.SIGNUP, func(c *routing.Context) error {
		return conf.Signup(string(c.PostBody()), c)
		//return somefolder.somefunc(string(c.PostBody()), c)
	})
	r.Post(constants.DS+constants.API+constants.DS+constants.CREDIT, func(c *routing.Context) error {
		return conf.Credit(string(c.PostBody()), c)
		//return somefolder.somefunc(string(c.PostBody()), c)
	})
	r.Post(constants.DS+constants.API+constants.DS+constants.DEBIT, func(c *routing.Context) error {
		return conf.Debit(string(c.PostBody()), c)
		//return somefolder.somefunc(string(c.PostBody()), c)
	})
	r.Post(constants.DS+constants.API+constants.DS+constants.SEND, func(c *routing.Context) error {
		return conf.Send(string(c.PostBody()), c)
		//return somefolder.somefunc(string(c.PostBody()), c)
	})
	r.Post(constants.DS+constants.API+constants.DS+constants.REFUND, func(c *routing.Context) error {
		return conf.RefundResponse(string(c.PostBody()), c)
		//return somefolder.somefunc(string(c.PostBody()), c)
	})
	r.Get(constants.DS+constants.API+constants.DS+constants.HISTORY, func(c *routing.Context) error {
		conf.History(c)
		return nil
	})
}
