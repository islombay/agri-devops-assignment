package service

import (
	"time"

	"github.com/islombay/agri-devops-assignment/internal/model"
)

type AgricultureService struct {
	farmers   []model.Farmer
	suppliers []model.Supplier
	products  []model.Product
	orders    []model.Order
}

func NewAgricultureService() *AgricultureService {
	return &AgricultureService{
		farmers: []model.Farmer{
			{"F1", "John Doe", "California", "Wheat", 50.5},
			{"F2", "Jane Smith", "Texas", "Corn", 120.0},
			{"F3", "Bob Johnson", "Iowa", "Soybean", 80.0},
			{"F4", "Alice Brown", "Florida", "Oranges", 30.0},
		},
		suppliers: []model.Supplier{
			{"S1", "AgriCorp", "Fertilizers", "contact@agricorp.com"},
			{"S2", "FarmMachinery Inc", "Machinery", "sales@farmmachinery.com"},
			{"S3", "SeedMaster", "Seeds", "support@seedmaster.com"},
		},
		products: []model.Product{
			{"P1", "Nitrogen Fertilizer", "Fertilizers", 50.0, "S1"},
			{"P2", "Harvester 3000", "Machinery", 150000.0, "S2"},
			{"P3", "Corn Seeds", "Seeds", 200.0, "S3"},
		},
		orders: []model.Order{
			{"O1", "F1", "P1", 10, 500.0, time.Now().AddDate(0, 0, -5), "DELIVERED"},
			{"O2", "F2", "P2", 1, 150000.0, time.Now().AddDate(0, 0, -2), "PROCESSING"},
			{"O3", "F3", "P3", 5, 1000.0, time.Now().AddDate(0, 0, -1), "SHIPPED"},
		},
	}
}

func (s *AgricultureService) GetFarmers() []model.Farmer     { return s.farmers }
func (s *AgricultureService) GetSuppliers() []model.Supplier { return s.suppliers }
func (s *AgricultureService) GetProducts() []model.Product   { return s.products }
func (s *AgricultureService) GetOrders() []model.Order       { return s.orders }
