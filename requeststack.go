package main

import (
	"net/http"
)

type RequestStack []*Request

func (rs *RequestStack) add(r *Request) {
	*rs = append(*rs, r)
}

type Request struct {
	Details *http.Request
	Body    string
}
