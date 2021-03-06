package lcd

func NewThirdLineLeftSegment() Segment{
    return &SideSegment{
        map[int]bool {
            1: true,
            3: true,
            4: true,
            5: true,
            7: true,
            9: true,
        },
    }
}

func NewThirdLineCenterSegment() Segment {
    return &CenterSegment{
        map[int]bool {
            1: true,
            4: true,
            7: true,
            9: true,
        },
    }
}

func NewThirdLineRightSegment() Segment {
    return &SideSegment{
        map[int]bool{
            2: true,
        },
    }
}
