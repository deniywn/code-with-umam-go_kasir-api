package repositories

import (
	"database/sql"
	models "kasir-api/model"
	"log"
)

// type ReportRepository interface {
// 	GetTodaySalesSummary() (*models.SalesSummary, error)
// 	GetSalesSummaryByDateRange(startDate, endDate string) (*models.SalesSummary, error)
// }

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetTodaySalesSummary() (*models.SalesSummary, error) {
	query := `
      SELECT
               COALESCE(SUM(total_price0) AS total_sales,
               COALESCE(COUNT(*0) AS total_transactions
      FROM
               "transactions"
      WHERE
               DATE(transaction_date) = CURRENT_DATE`
	row := r.db.QueryRow(query)

	var summary models.SalesSummary
	err := row.Scan(&summary.TotalRevenue, &summary.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func (r *ReportRepository) GetSalesSummaryByDateRange(startDate, endDate string) (*models.SalesSummary, error) {
	summary := &models.SalesSummary{}

	// Query 1: Mendapatkan Total Revenue dan Total Transaksi dalam rentang tanggal
	err := r.db.QueryRow(`
        SELECT
            COALESCE(SUM(total_amount), 0) as total_revenue,
            COUNT(*) as total_transaksi
        FROM transactions
        WHERE DATE(created_at) BETWEEN $1 AND $2
    `, startDate, endDate).Scan(&summary.TotalRevenue, &summary.TotalTransaksi)

	if err != nil {
		log.Printf("Error querying sales summary by date range: %v", err)
		return nil, err
	}

	// Query 2: Mendapatkan Produk Terlaris dalam rentang tanggal
	err = r.db.QueryRow(`
        SELECT
            p.name,
            SUM(td.quantity) as total_quantity
        FROM transaction_details td
        JOIN products p ON td.product_id = p.id
        JOIN transactions t ON td.transaction_id = t.id
        WHERE DATE(t.created_at) BETWEEN $1 AND $2
        GROUP BY p.name
        ORDER BY total_quantity DESC
        LIMIT 1
    `, startDate, endDate).Scan(&summary.ProdukTerlaris.Nama, &summary.ProdukTerlaris.QtyTerjual)

	// Jika tidak ada produk yang terjual (sql.ErrNoRows), biarkan data produk terlaris kosong.
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error querying best selling product by date range: %v", err)
		return nil, err
	}

	return summary, nil
}
