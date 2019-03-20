package models


func seed() {
	tenant := Tenant{
		Name: "电竞大师",
	}
	DB.Create(&tenant)
}