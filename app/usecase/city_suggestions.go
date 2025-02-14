package usecase

import (
	"context"
	"math"
	"sort"
	"strconv"
	"strings"
)

func (uc *Usecase) GetCitySuggestions(ctx context.Context, req RequestScoredCity) (resp []ResponseScoredCity, err error) {
	cities := uc.Repository.DataSource.Search(req.Query)

	lat, lon := parseCoordinates(req.LatStr, req.LonStr)
	for _, city := range cities {
		prefixScore := calculatePrefixScore(req.Query, city.Name)
		levenshteinScore := calculateLevenshteinScore(req.Query, city.Name)
		geoScore := 0.0

		if lat != 0 && lon != 0 {
			distance := haversine(lat, lon, city.Latitude, city.Longitude)
			geoScore = 1 - (distance / 5000.0) // Normalize (Assume 5000km max)
		}

		finalScore := (prefixScore * 0.5) + (levenshteinScore * 0.3) + (geoScore * 0.2)
		finalScore = math.Round(finalScore*10) / 10

		resp = append(resp, ResponseScoredCity{
			Name:      city.Name,
			Latitude:  city.Latitude,
			Longitude: city.Longitude,
			Score:     finalScore,
		})
	}

	// Sort by highest score
	sort.Slice(resp, func(i, j int) bool {
		return resp[i].Score > resp[j].Score
	})
	return
}

// Prefix Match Score
func calculatePrefixScore(query, cityName string) float64 {
	query = strings.ToLower(query)
	cityName = strings.ToLower(cityName)

	if strings.HasPrefix(cityName, query) {
		return float64(len(query)) / float64(len(cityName))
	}
	return 0.0
}

// Levenshtein Distance Score
func calculateLevenshteinScore(query, cityName string) float64 {
	dist := levenshtein(query, cityName)
	maxLen := math.Max(float64(len(query)), float64(len(cityName)))
	return 1 - (float64(dist) / maxLen)
}

// Parse Latitude & Longitude
func parseCoordinates(latStr, lonStr string) (float64, float64) {
	lat, _ := strconv.ParseFloat(latStr, 64)
	lon, _ := strconv.ParseFloat(lonStr, 64)
	return lat, lon
}

// Haversine formula to calculate distance in km
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Earth radius in KM
	dLat := (lat2 - lat1) * (math.Pi / 180)
	dLon := (lon2 - lon1) * (math.Pi / 180)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*(math.Pi/180))*math.Cos(lat2*(math.Pi/180))*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

// Levenshtein distance function
func levenshtein(s1, s2 string) int {
	row1 := make([]int, len(s2)+1)
	row2 := make([]int, len(s2)+1)

	for i := range row1 {
		row1[i] = i
	}

	for i, c1 := range s1 {
		row2[0] = i + 1

		for j, c2 := range s2 {
			insertions := row1[j+1] + 1
			deletions := row2[j] + 1
			substitutions := row1[j]
			if c1 != c2 {
				substitutions++
			}

			row2[j+1] = min(insertions, deletions, substitutions)
		}
		row1, row2 = row2, row1
	}
	return row1[len(s2)]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
