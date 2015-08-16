package lcd

func NewSecondLineLeftSegment() Segment {
	return &SideSegment{
		map[int]bool{
			1: true,
			2: true,
			3: true,
		},
	}
}

func NewSecondLineCenterSegment() Segment {
	return &CenterSegment{
		map[int]bool{
			2: true,
			3: true,
			4: true,
			5: true,
		},
	}
}
