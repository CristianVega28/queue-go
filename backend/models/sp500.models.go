package models

type (
	SP500 struct {
		Exchange      string
		Symbol        string
		Shortname     string
		Longname      string
		Sector        string
		Industry      string
		CurrentPrice  float32
		MarketCap     float32
		Ebitda        float32
		RevenueGrowth float32
		Model
	}
)

func (sp SP500) GetModel() Model {
	return sp.Model
}
