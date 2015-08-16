package lcd

func NewThirdLineCenterSegment() Segment {
    return &CenterSegment{
        map[int]bool {
            1:true,
            4:true,
        },
    }
}
