package services

type HealthService struct{}

func (s *HealthService) GetStatus() string {
	return "OK"
}
