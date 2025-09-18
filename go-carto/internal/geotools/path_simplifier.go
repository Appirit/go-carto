package geotools

// enelève les points sans distance (possiblement car le HR varie dans le temps)
func removeDuplicate(points []Position, intervals []IntervalDistance) ([]Position, []IntervalDistance) {
	n := 0
	for iInt, seg := range intervals {
		if seg != 0 {
			if iInt != n {
				points[n+1] = points[iInt+1]
				intervals[n] = intervals[iInt]
			}
			n++
		}
	}
	points[n-1] = points[len(points)-1] // le dernier subsite
	return points[:n], intervals[:n]
}
