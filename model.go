package main

type Stats struct {
	Min, Max, Sum, Avg float64
	Count              int64
}

type StatsInt32 struct {
	Min, Max int32
	Sum      int64
	Avg      float64
	Count    int64
}
