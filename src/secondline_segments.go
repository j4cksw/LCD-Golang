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
