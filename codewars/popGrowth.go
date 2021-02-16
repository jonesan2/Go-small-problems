package kata

func NbYear(p0 int, percent float64, aug int, p int) (years int) {
	curPop := p0
	for {
		growthFromPerc := int(float64(curPop) * percent / 100.0)
		curPop = curPop + growthFromPerc + aug
		years = years + 1
		if curPop >= p {
			return years
		}
	}
}
