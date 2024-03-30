package app

import "testing"

func BenchmarkExecute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Execute("../../data/weather_data.csv")
	}

}
