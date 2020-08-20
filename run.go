package hybridsort

// run is a chunk of data to be sorted and merged
type run struct {
	Start int
	End   int
	Seq   int
}

// Len returns run's length
func (r run) Len() int { return 1 + r.End - r.Start }
