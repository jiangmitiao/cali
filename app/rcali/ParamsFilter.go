package rcali

import (
	"github.com/revel/revel"
	"net/url"
)

func copyValues(dst, src url.Values) {
	for k, vs := range src {
		for _, value := range vs {
			dst.Add(k, value)
		}
	}
}

//query info to request.form,form exist first
func QueryParamsFilter(c *revel.Controller, fc []revel.Filter) {
	newValues, e := url.ParseQuery(c.Request.URL.RawQuery)
	if e == nil {
		if c.Request.Form==nil {
			c.Request.Form = make(url.Values)
		}
		copyValues(c.Request.Form,newValues)
	}
	fc[0](c, fc[1:])
}
