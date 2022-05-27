package price

type Product struct {
	Name   string
	Price  float64
	Amount int
}

type Service struct {
	Name         string
	Price        float64
	WorkedMinute int
}

type Maintenance struct {
	Name  string
	Price float64
}

func ProductsPrice(products []Product) float64 {
	totalValue := 0.0
	for _, product := range products {
		totalValue += product.Price * float64(product.Amount)
	}
	return totalValue
}

func ServicesPrice(services []Service) float64 {
	totalValue := 0.0
	for _, service := range services {
		if service.WorkedMinute < 30 {
			totalValue += service.Price * 0.5
		} else {
			hoursWorked := float64(service.WorkedMinute) / float64(60)
			totalValue += service.Price * hoursWorked
		}
	}
	return totalValue
}

func MaintenancePrice(maintenances []Maintenance) float64 {
	totalValue := 0.0
	for _, maintenance := range maintenances {
		totalValue += maintenance.Price
	}
	return float64(totalValue)
}
