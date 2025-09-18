package geotools

import (
	"fmt"
)

func CompareMethods(name string, points LineString) {
	fmt.Printf("*** %s ***\n", name)
	// sements := BuildRouteSegments(&points)
	for i := 1; i < 15; i += 2 {
		met := NewElevationMetrics(SmoothMovingAverage(points, i))
		fmt.Printf("SmoothMovingAverage - Moving average simple : (ech. %d) %s\n", i, met.ToString())
	}

	met := NewElevationMetrics(SmoothMedian(points, 10))
	fmt.Printf("SmoothMedian - Median filter pour supprimer spikes :  %s\n", met.ToString())

	met = NewElevationMetrics(SmoothByDistance(points, 10))
	fmt.Printf("SmoothByDistance - Lissage par fenêtre en mètres (utile si échantillonnage irrégulier) :  %s\n", met.ToString())

}
func Measure(name string, points LineString) {

	intervalle := calculateEachDistance(points)
	points, intervalle = removeDuplicate(points, intervalle)

	fmt.Printf("*** DISTANCE de '%s' ***\n", name)
	// sements := BuildRouteSegments(&points)
	// for i := 1; i < 15; i += 2 {
	// 	met := NewElevationMetrics(SmoothMovingAverage(points, i))
	// 	fmt.Printf("SmoothMovingAverage - Moving average simple : (ech. %d) %s\n", i, met.ToString())
	// }
	met := NewElevationMetrics(SmoothMovingAverage(points, amplitudeLissage))
	fmt.Printf("SmoothMovingAverage - Moving average simple : (ech. %d) %s\n", amplitudeLissage, met.ToString())

	// met = NewElevationMetrics(SmoothMedian(points, 10))
	// fmt.Printf("SmoothMedian - Median filter pour supprimer spikes :  %s\n", met.ToString())

	// met = NewElevationMetrics(SmoothByDistance(points, 10))
	// fmt.Printf("SmoothByDistance - Lissage par fenêtre en mètres (utile si échantillonnage irrégulier) :  %s\n", met.ToString())
	fmt.Printf("*** PAUSE de '%s' ***\n", name)

	var roulants = rollingSegments(points, intervalle, 2, pauseMini)
	var cum float64
	for i, r := range roulants {
		d := duree(points, r.Start, r.End-1)
		cum += d
		fmt.Printf("	%d. [%d, %d] %s\n", i+1, r.Start, r.End, ToString(d))
	}
	fmt.Printf("	pour une durée roulante de %s\n", ToString(cum))
	// for sec := 10; sec < 3*60; sec += 10 {
	// 	var roulants = rollingSegments(points, intervalle, 2, Duree(sec))
	// 	var cum float64
	// 	for _, r := range roulants {
	// 		d := duree(points, r.Start, r.End-1)
	// 		cum += d
	// 	}
	// 	fmt.Printf("	pause de %s = durée de %s\n", ToString(Duree(sec)), ToString(cum))

	// }
}
