package models

type CityModel struct {
	Id                  int      `json:"id"`
	Name                string   `json:"name"`
	Country             string   `json:"country"`
	Population          int64    `json:"population"`
	NoisyPollutionLevel string   `json:"noisy_pollution_level"`
	NeighbourhoodCities []string `json:"neighbourhood_cities"`
}
