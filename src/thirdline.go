package lcd

func NewThirdLine() Line {
    return Line{
        []Segment{
            NewThirdLineLeftSegment(),
            NewThirdLineCenterSegment(),
            NewThirdLineRightSegment(),
        },
    }
}
