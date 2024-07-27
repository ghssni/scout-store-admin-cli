package handler

import (
	"scout-store-admin-cli/repository/product"
	"scout-store-admin-cli/repository/sales"
	"scout-store-admin-cli/repository/staff"
)

type Handler struct {
	ProductRepo *product.Repository
	StaffRepo   *staff.Repository
	SalesRepo   *sales.Repository
}

func NewHandler(productRepo *product.Repository, staffRepo *staff.Repository, salesRepo *sales.Repository) *Handler {
	return &Handler{
		ProductRepo: productRepo,
		StaffRepo:   staffRepo,
		SalesRepo:   salesRepo,
	}
}
