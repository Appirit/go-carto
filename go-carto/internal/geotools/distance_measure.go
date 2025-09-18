package geotools

import (
	"math"
)

// func (curseg *RouteSegment) add(seg *RouteSegment) {
// 	curseg.ID = seg.ID // dernier id
// 	curseg.Distance_m += seg.Distance_m
// 	curseg.Duration_sec += seg.Duration_sec
// }

// Générateur de Segments de route
func calculateEachDistance(points LineString) []IntervalDistance {
	n := len(points)
	if n < 2 {
		return nil
	}
	res := make([]IntervalDistance, n)

	for i, pt := range points {
		if i > 0 {
			dx := haversine(pt[lat], pt[lon], points[i-1][lat], points[i-1][lon])
			dy := float64(0)
			if len(points) >= ele_m {
				dy = points[i][ele_m] - points[i-1][ele_m]
			}
			if dy != 0 {
				// distance 3D
				math.Sqrt(dx*dx + dy*dy)
			} else {
				res[i-1] = dx
			}
		}
	}

	return res
}

// Fonction pour calculer la distance avec la formule de Haversine en mètres
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 63727954.77598 // rayon R = 6372.795477598 km (rayon moyen quadrique)
	// const R = 6371000.0
	φ1 := lat1 * math.Pi / 180
	φ2 := lat2 * math.Pi / 180
	dφ := (lat2 - lat1) * math.Pi / 180
	dλ := (lon2 - lon1) * math.Pi / 180
	a := math.Sin(dφ/2)*math.Sin(dφ/2) + math.Cos(φ1)*math.Cos(φ2)*math.Sin(dλ/2)*math.Sin(dλ/2)
	return 2 * R * math.Asin(math.Min(1, math.Sqrt(a)))
}

// func newSegment(previousPt, currentPt Position) RouteSegment {
// 	seg := RouteSegment{
// 		Distance_m:   haversine(currentPt[lat], currentPt[lon], previousPt[lat], previousPt[lon]),
// 		Duration_sec: currentPt[time_sec] - previousPt[time_sec],
// 		Ascent:       currentPt[ele_m] - previousPt[ele_m],
// 	}

// 	if seg.Ascent != 0 {
// 		// correction avec la pente
// 		seg.Distance_m = math.Sqrt(math.Pow(seg.Distance_m, 2) + math.Pow(seg.Ascent, 2))
// 	}
// 	return seg
// }

// func (segments *RouteSegments) DistanceFromGeometry() float64 {
// 	dist := float64(0)
// 	for _, seg := range *segments {
// 		dist += seg.Distance_m
// 	}
// 	return math.Round(dist)
// }
