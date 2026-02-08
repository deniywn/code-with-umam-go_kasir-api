package services

import (
	models "kasir-api/model"
	"kasir-api/repositories"
)

// type ReportService interface {
// 	GenerateSalesSummaryByDateRange(startDate, endDate string) (*models.SalesSummary, error)
// }

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetTodaySalesSummary() (*models.SalesSummary, error) {
	return s.repo.GetTodaySalesSummary()
}

func (s *ReportService) GenerateSalesSummaryByDateRange(startDate, endDate string) (*models.SalesSummary, error) {
	return s.repo.GetSalesSummaryByDateRange(startDate, endDate)
}
