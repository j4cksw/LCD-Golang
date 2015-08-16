package lcd

func NewThirdLine() Line {
    return &ConcreteLine{
        "|_||_|",
        []Segment{
            NewThirdLineLeftSegment(),
            NewThirdLineCenterSegment(),
            NewThirdLineRightSegment(),
        },
    }
}
