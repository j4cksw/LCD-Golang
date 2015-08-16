package lcd

func NewThirdLine() Line {
    return Line{
        "|_||_|",
        []Segment{
            NewThirdLineLeftSegment(),
            NewThirdLineCenterSegment(),
            NewThirdLineRightSegment(),
        },
    }
}
