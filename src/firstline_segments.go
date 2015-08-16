package lcd

func NewFirstLineCenterSegment() Segment{
    return &CenterSegment{
        map[int]bool{
            1: true,
            4: true,
        },
    }
}

func NewFirstLineLeftSegment() Segment{
    return &AlwaysOffSegment{}
}

func NewFirstLineRightSegment() Segment{
    return &AlwaysOffSegment{}
}
