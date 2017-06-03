package rcali

import (
	"github.com/revel/revel"
	"net/http"
)

type IMGJPG []byte

func (i IMGJPG) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusOK, "image/jpg")
	resp.Out.Write(i)
}

type FILE []byte

func (i FILE) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusOK, "application/octet-stream")
	resp.Out.Write(i)
}
