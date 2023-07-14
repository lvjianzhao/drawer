package response

type CustomerData struct {
	Total   TotalData   `json:"total"`
	Analyse AnalyseData `json:"analyse"`
}

type TotalData struct {
	Customer int64 `json:"customer"`
	Env      int64 `json:"env"`
	License  int64 `json:"license"`
}

type AnalyseData struct {
	CustomerType     []AnalyseItem `json:"customerType"`
	TechnicalSupport []AnalyseItem `json:"technicalSupport"`
	MaintenanceType  []AnalyseItem `json:"maintenanceType"`
	Env              []AnalyseItem `json:"env"`
	License          []AnalyseItem `json:"license"`
}

type AnalyseItem struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
}

type CustomerType struct {
	Type  uint `gorm:"column:type"`
	Count int  `gorm:"column:count"`
}

type TechnicalSupport struct {
	Status uint `gorm:"column:service_status"`
	Count  int  `gorm:"column:count"`
}

type Env struct {
	Env   string `gorm:"column:env"`
	Count int    `gorm:"column:count"`
}

type License struct {
	Active uint `gorm:"column:active"`
	Count  int  `gorm:"column:count"`
}

type MaintenanceType struct {
	Type  uint `gorm:"column:maintenance_type"`
	Count int  `gorm:"column:count"`
}
