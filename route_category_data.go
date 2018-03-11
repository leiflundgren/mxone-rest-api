package main

type RouteCategoryData struct {
	Rou     string `json:"route"`
	Cust    string `json:"cust"`
	selRaw  string
	Trm     string `json:"trm"`
	servRaw string
	Nodg    string `json:"nodg"`
	Dist    string `json:"dist"`
	Disl    string `json:"disl"`
	trafRaw string
	sigRaw  string
	bcapRaw string
}

func (routeCategory *RouteCategoryData) ReadFromFile() []RouteCategoryData {

}
