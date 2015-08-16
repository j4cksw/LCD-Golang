package lcd

func NewThirdLine() Line {
    return Line{
        "|_|",
        []Segment{
            NewThirdLineLeftSegment(),
            NewThirdLineCenterSegment(),
            NewThirdLineRightSegment(),
        },
    }
}
