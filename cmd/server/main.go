package main

import (
	"fmt"

	"github.com/UGRORF/go-musthave-metrics-tpl/internal/handler"
	"github.com/UGRORF/go-musthave-metrics-tpl/internal/repository"
	"github.com/UGRORF/go-musthave-metrics-tpl/internal/service"
)

func main() {
	metricStorage := repository.NewMemStorage()
	metricService := service.NewMetricService(metricStorage)
	metricHandler := handler.NewMetricHandler(metricService)

	server := handler.NewServer(":8080", metricHandler)
	if err := server.Start(); err != nil {
		fmt.Printf("Server failed: %v", err)
	}
}
