package service

type MetricRepository interface {
	UpdateGauge(name string, value float64)
	UpdateCounter(name string, value int64)
}

type MetricService struct {
	repo MetricRepository
}

func NewMetricService(repo MetricRepository) *MetricService {
	return &MetricService{
		repo: repo,
	}
}

func (s *MetricService) UpdateGauge(name string, value float64) error {
	s.repo.UpdateGauge(name, value)
	return nil
}

func (s *MetricService) UpdateCounter(name string, value int64) error {
	s.repo.UpdateCounter(name, value)
	return nil
}
