package handler

import (
	"net/http"
)

func routerInit(metricHandler *MetricHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/update/", middleware(http.HandlerFunc(metricHandler.updateHandler)))
	return mux
}
