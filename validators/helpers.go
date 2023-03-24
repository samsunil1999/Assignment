package validators

import "github.com/thedevsaddam/govalidator"

func getRulesForAddProduct() govalidator.MapData {
	rules := govalidator.MapData{
		"name":     []string{"required"},
		"price":    []string{"required", "float"},
		"category": []string{"required"},
		"quantity": []string{"required", "numeric"},
	}
	return rules
}

func getRulesForUpdateProduct() govalidator.MapData {
	rules := govalidator.MapData{
		"price":    []string{"float"},
		"quantity": []string{"numeric"},
	}
	return rules
}

func getRulesForOrderedProductData() govalidator.MapData {
	rules := govalidator.MapData{
		"product_id":       []string{"required"},
		"product_quantity": []string{"required", "numeric"},
	}
	return rules
}
