package main

import (
	"fmt"
	"mapSharedRead1/loader"
	"sort"
)

func sortBy[Key comparable, Val any](data map[Key]Val, lessThan func(r1, r2 Val) bool) chan []Key {
	ch := make(chan []Key)

	go func() {
		keys := make([]Key, 0, len(data))

		for ident := range data {
			keys = append(keys, ident)
		}

		sort.Slice(keys, func(i, j int) bool {
			return lessThan(data[keys[i]], data[keys[j]])
		})

		ch <- keys
		close(ch)
	}()

	return ch
}

func sortIdentsByElevation(airports map[string]loader.Record) chan []string {
	return sortBy(airports, func(r1, r2 loader.Record) bool {
		return r1.ElevationFt < r2.ElevationFt
	})
}

func sortIdentsByLong(airports map[string]loader.Record) chan []string {
	return sortBy(airports, func(r1, r2 loader.Record) bool {
		return r1.Longitude < r2.Longitude
	})
}

func main() {
	dataMap := loader.GetAll()
	fmt.Printf("loaded %v airports\n", len(dataMap))

	sortedByLongCh := sortIdentsByLong(dataMap)
	sortedByElevationCh := sortIdentsByElevation(dataMap)

	complete := 2
	for complete > 0 {
		complete--

		select {
		case sElev := <-sortedByElevationCh:
			max := sElev[len(sElev)-1]
			min := sElev[0]

			fmt.Printf("lowest airport: %v\n", dataMap[min])
			fmt.Printf("highest airport: %v\n", dataMap[max])

		case sLong := <-sortedByLongCh:
			max := sLong[len(sLong)-1]
			min := sLong[0]

			fmt.Printf("western airport: %v\n", dataMap[min])
			fmt.Printf("eastern airport: %v\n", dataMap[max])
		}
	}
}
