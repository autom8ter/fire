package webtask

import (
	"github.com/autom8ter/api/driver"
	"net/http"
)

type WebTask struct {
	driver.Callbacker
	driver.ErrorHandler
	http.RoundTripper
}

func NewWebTask(c driver.CallbackFunc, e driver.ErrorFunc, h driver.RoundTripperFunc) *WebTask {
	return &WebTask{
		Callbacker:   c,
		ErrorHandler: e,
		RoundTripper: h,
	}
}

func (w *WebTask) HandlerFunc() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		resp, err := w.RoundTrip(req)
		if err != nil {
			w.HandleError(writer, err)
			return
		}
		err = w.Callback(writer, resp.Body)
		if err != nil {
			w.HandleError(writer, err)
			return
		}
	}
}
