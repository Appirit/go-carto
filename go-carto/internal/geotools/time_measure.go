package geotools

import (
	"fmt"
	"math"
)

type Segment struct{ Start, End int } // couple d'index : start stop

func duree(coords []Position, start, end int) Duree {
	t1 := coords[start][time_sec]
	t2 := coords[end][time_sec]
	// fmt.Println(t2 - t1)
	return t2 - t1
}

func ToString(duree Duree) string {
	var values []float64
	var unites []string
	if duree != 1 {
		// sec
		values = append(values, math.Mod(duree, 60))
		unites = append(unites, "\"")
		duree /= 60
	}
	if duree >= 1 {
		// mn
		values = append(values, math.Mod(duree, 60))
		unites = append(unites, "'")
		duree /= 60
	}
	if duree >= 1 {
		// hr
		values = append(values, math.Mod(duree, 24))
		unites = append(unites, "h")
		duree /= 24
	}
	if duree >= 1 {
		// jours
		values = append(values, math.Mod(duree, 31))
		unites = append(unites, "j")
		duree /= 31
	}
	if duree >= 1 {
		// mois
		values = append(values, duree)
		unites = append(unites, "mois")
	}
	res := ""
	for i := len(values) - 1; i >= 0; i-- {
		if i != len(values)-1 {
			res += " "
		}
		res += fmt.Sprintf("%d%s", int(values[i]), unites[i])
	}
	return res
}

func rollingSegments(coords []Position, intervals []IntervalDistance, minSpeed, pauseMin Duree) []Segment {
	n := len(coords)
	if n < 2 {
		return nil
	}
	var segs []Segment
	var moveStart, pauseStart int
	var inMove, pendingMove bool

	for i := 1; i < n; i++ {
		t := duree(coords, i-1, i)
		if t <= 0 {
			continue
		}
		speed := intervals[i-1] / t // m/s

		if speed >= minSpeed {
			// 2 points en mouvement
			if !inMove || i == 1 {
				// on passe en mouvement
				inMove = true
				if !pendingMove {

					moveStart = i - 1
					pendingMove = true
				}
			}
		} else {
			// pause
			if inMove || i == 1 {
				// on passe en pause
				inMove = false
				pauseStart = i
			}

			if duree(coords, pauseStart, i) >= pauseMin && pendingMove {
				// si la pause dure assez longtemps → on clôture le segment
				segs = append(segs, Segment{Start: moveStart, End: pauseStart})
				pendingMove = false
			}
			// on laisse passer les pauses
		}
	}
	if pendingMove {
		segs = append(segs, Segment{Start: moveStart, End: n})
	}
	return segs
}

// func simplifyNoPause(points LineString, intervals []IntervalDistance, speedMini_ms float64) (newgeometry LineString) {
// 	pauseStart := points[0]
// 	memdist := float64(0)
// 	n := 0
// 	for i, seg := range intervals {
// 		// segment courant
// 		v := seg / (points[i][time_sec] - points[i-1][time_sec])

// 		// calcule de la pause
// 		pause := newSegment(pauseStart, pt)
// 		pause_v_ms := pause.Distance_m / pause.Duration_sec
// 		if pause_v_ms < speedMini_ms /* vitesse de la pause*/ && v < 2*speedMini_ms /*vitesse du segment*/ {
// 			n++
// 			// point ignoré
// 		} else {
// 			if n != 0 {
// 				log.Printf("%d - pause de %.1fm à %.1fm - %dpt, durant %.0fs (%d)\n", seg.ID, memdist, total.Distance_m, n, pause.Duration_sec, int(pt[time_sec]-(*points)[0][time_sec]))
// 			}
// 			memdist = total.Distance_m
// 			n = 0
// 			pauseStart = pt // pas de pause ici, la suite ?
// 			newgeometry = append(newgeometry, pt)

// 		}
// 	}
// 	log.Printf("fin à %.1fm\n", total.Distance_m)
// 	return
// }

// altitude selon le type de géométrie
// func getRouteTime(feature *geojson.Feature, met *RouteMetrics) {
// 	switch feature.Geometry.Type {

// 	case geojson.GeometryLineString:
// 		elevationFromGeometry(met, feature.Geometry.LineString)

// 	case geojson.GeometryMultiLineString:
// 		if len(feature.Geometry.MultiLineString) > 0 && len(feature.Geometry.MultiLineString[0]) > 0 && len(feature.Geometry.MultiLineString[0][0]) > 0 {
// 			for _, geometry := range feature.Geometry.MultiLineString {
// 				elevationFromGeometry(met, geometry)
// 			}
// 		}
// 	}
// }
