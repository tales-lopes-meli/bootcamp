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

func ProductsPrice(products []Product, c chan float64) float64 {
	totalValue := 0.0
	for _, product := range products {
		totalValue += product.Price * float64(product.Amount)
	}
	c <- totalValue
	return totalValue
}

func ServicesPrice(services []Service, c chan float64) float64 {
	totalValue := 0.0
	for _, service := range services {
		if service.WorkedMinute < 30 {
			totalValue += service.Price * 0.5
		} else {
			hoursWorked := float64(service.WorkedMinute) / float64(60)
			totalValue += service.Price * hoursWorked
		}
	}
	c <- totalValue
	return totalValue
}

func MaintenancePrice(maintenances []Maintenance, c chan float64) float64 {
	totalValue := 0.0
	for _, maintenance := range maintenances {
		totalValue += maintenance.Price
	}
	c <- totalValue
	return float64(totalValue)
}
