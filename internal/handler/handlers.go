package handler

import (
	"net/http"
	"strconv"
	"strings"

	models "github.com/UGRORF/go-musthave-metrics-tpl/internal/model"
	"github.com/UGRORF/go-musthave-metrics-tpl/internal/service"
)

type MetricHandler struct {
	metricService *service.MetricService
}

func NewMetricHandler(metricService *service.MetricService) *MetricHandler {
	return &MetricHandler{metricService: metricService}
}

func (s *MetricHandler) updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}
	path := strings.TrimPrefix(r.URL.Path, "/update/")
	partsOfPath := strings.Split(path, "/")

	if len(partsOfPath) < 3 || partsOfPath[1] == "" {
		http.Error(w, "Wrong request. Metric's name is missing", http.StatusNotFound)
		return
	}
	if len(partsOfPath) > 3 {
		http.Error(w, "Wrong request. Invalid URL format", http.StatusBadRequest)
		return
	}

	metricType := partsOfPath[0]
	metricName := partsOfPath[1]
	metricValue := partsOfPath[2]

	if metricType == models.Counter {
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			http.Error(w, "Invalid value", http.StatusBadRequest)
			return
		}
		if err = s.metricService.UpdateCounter(metricName, value); err != nil {
			http.Error(w, "Invalid value", http.StatusBadRequest)
			return
		}
	} else if metricType == models.Gauge {
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			http.Error(w, "Invalid value", http.StatusBadRequest)
			return
		}
		if err = s.metricService.UpdateGauge(metricName, value); err != nil {
			http.Error(w, "Invalid value", http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "Type of metric is wrong", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome"))
}
