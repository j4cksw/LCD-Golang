package lcd

func NewThirdLineLeftSegment() Segment{
    return &SideSegment{
        map[int]bool {
            1: true,
            3: true,
            4: true,
            5: true,
        },
    }
}
