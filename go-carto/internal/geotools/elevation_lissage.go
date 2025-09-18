package geotools

import (
	"math"
	"sort"
)

// revient à utiliser SmoothMovingAverage(xxx, 1 ou 0)
// func NoCorrection(points LineString) []float64 {
// 	n := len(points)
// 	out := make([]float64, n)
// 	for i := 0; i < n; i++ {
// 		out[i] = points[i][ele_m]
// 	}
// 	return out
// }

// Moving average simple
// alt: altitudes GPS
// window: nombre d'échantillons pairs/impairs (ex: 5, 11)
// renvoie altitudes lissées
func SmoothMovingAverage(points LineString, window int) []float64 {
	n := len(points)
	out := make([]float64, n)
	if n == 0 || window <= 1 {
		for i := 0; i < n; i++ {
			out[i] = points[i][ele_m]
		}
		return out
	}
	if window > n {
		window = n
	}
	half := window / 2
	var sum float64
	// initial sum
	for i := 0; i < window && i < n; i++ {
		sum += points[i][ele_m]
	}
	for i := 0; i < n; i++ {
		start := i - half
		if start < 0 {
			start = 0
		}
		end := i + half
		if end >= n {
			end = n - 1
		}
		// recompute sum for small windows (simple and safe)
		sum = 0
		for k := start; k <= end; k++ {
			sum += points[k][ele_m]
		}
		out[i] = sum / float64(end-start+1)
	}
	return out
}

// Median filter pour supprimer spikes
func SmoothMedian(points LineString, window int) []float64 {
	n := len(points)
	out := make([]float64, n)
	if n == 0 || window <= 1 {
		for i := 0; i < n; i++ {
			out[i] = points[i][ele_m]
		}
		return out
	}
	if window > n {
		window = n
	}
	half := window / 2
	buf := make([]float64, 0, window)
	for i := 0; i < n; i++ {
		buf = buf[:0]
		start := i - half
		if start < 0 {
			start = 0
		}
		end := i + half
		if end >= n {
			end = n - 1
		}
		for k := start; k <= end; k++ {
			buf = append(buf, points[k][ele_m])
		}
		sort.Float64s(buf)
		out[i] = buf[len(buf)/2]
	}
	return out
}

// Lissage par fenêtre en mètres (utile si échantillonnage irrégulier)
// lat/lon en degrés. windowMeters ex: 20.0 ou 50.0
func SmoothByDistance(points LineString, windowMeters float64) []float64 {
	n := len(points)
	out := make([]float64, n)
	for i := 0; i < n; i++ {
		var sum float64
		var wcount int
		for k := 0; k < n; k++ {
			if math.Abs(float64(k-i)) > 5000 { // coupe rapide si indices très éloignés
				continue
			}
			if haversine(points[i][lat], points[i][lon], points[k][lat], points[k][lon]) <= windowMeters {
				sum += points[k][ele_m]
				wcount++
			}
		}
		if wcount == 0 {
			out[i] = points[i][ele_m]
		} else {
			out[i] = sum / float64(wcount)
		}
	}
	return out
}
