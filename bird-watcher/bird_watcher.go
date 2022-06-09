package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) (sum int) {
	for _, value := range birdsPerDay {
		sum += value
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) (sum int) {
	offset := 7 * (week - 1)
	return TotalBirdCount(birdsPerDay[offset : offset+7])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	i := 0
	for i < len(birdsPerDay) {
		birdsPerDay[i]++
		i += 2
	}
	return birdsPerDay
}
