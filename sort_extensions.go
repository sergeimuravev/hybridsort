package hybridsort

// WithMinRunSize assigns new min run size value.
func (sort Sort) WithMinRunSize(size uint8) Sort {
	iSize := int(size)
	if sort.maxRunSize > 0 && iSize > sort.maxRunSize {
		panic(size)
	}

	sort.minRunSize = iSize
	return sort
}

// WithMaxRunSize assigns new max run size value.
func (sort Sort) WithMaxRunSize(size uint8) Sort {
	iSize := int(size)
	if sort.minRunSize > 0 && iSize < sort.minRunSize {
		panic(size)
	}

	sort.maxRunSize = iSize
	return sort
}

// WithDegreeOfParallelism assigns new degree of parallelism value.
func (sort Sort) WithDegreeOfParallelism(dop uint8) Sort {
	if dop == 0 {
		panic(dop)
	}

	sort.dop = int(dop)
	return sort
}
