package lcd

func NewSecondLine() Line {
    return &ConcreteLine{
        "| || |",
        []Segment{
            NewSecondLineLeftSegment(),
            NewSecondLineCenterSegment(),
            NewSecondLineRightSegment(),
        },
    }
}
