package lcd

func NewSecondLine() Line {
    return Line{
        "| || |",
        []Segment{
            NewSecondLineLeftSegment(),
            NewSecondLineCenterSegment(),
            NewSecondLineRightSegment(),
        },
    }
}
