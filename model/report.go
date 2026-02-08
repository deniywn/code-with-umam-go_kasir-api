package models

// SalesSummary represents the daily sales summary.
type SalesSummary struct {
	TotalRevenue   float64        `json:"total_revenue"`
	TotalTransaksi int            `json:"total_transaksi"`
	ProdukTerlaris ProdukTerlaris `json:"produk_terlaris"`
}

// ProdukTerlaris represents the best-selling product.
type ProdukTerlaris struct {
	Nama       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}
