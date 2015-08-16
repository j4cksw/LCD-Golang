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
			0: true,
			1: true,
		},
	}
}

func NewSecondLineRightSegment() Segment {
	return &SideSegment{
		map[int]bool{
            5:true,
			6:true,
        },
	}
}
