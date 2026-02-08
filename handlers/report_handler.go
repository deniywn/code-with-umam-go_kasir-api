package handlers

import (
	"encoding/json"
	"kasir-api/services"
	"net/http"
	"time"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) GetTodaySalesSummary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	summary, err := h.service.GetTodaySalesSummary()
	if err != nil {
		// Jika ada error dari service, kirim response error 500
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(summary)
}

func (h *ReportHandler) GetSalesSummary(w http.ResponseWriter, r *http.Request) {
	// Ambil query parameter
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")

	// Default ke hari ini jika parameter tidak ada
	if startDateStr == "" {
		startDateStr = time.Now().Format("2006-01-02")
	}
	if endDateStr == "" {
		endDateStr = time.Now().Format("2006-01-02")
	}

	// TODO: Tambahkan validasi format tanggal jika diperlukan

	summary, err := h.service.GenerateSalesSummaryByDateRange(startDateStr, endDateStr)
	if err != nil {
		http.Error(w, "Gagal mengambil data ringkasan", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(summary)
}
