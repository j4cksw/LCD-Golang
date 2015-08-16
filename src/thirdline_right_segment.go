package lcd

func NewThirdLineRightSegment() Segment {
    return &SideSegment{
        map[int]bool{
            2: true,
        },
    }
}
