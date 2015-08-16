package lcd

func NewFirstLine() Line {
    return &ConcreteLine{
        " _  _ ",
        []Segment{
            NewFirstLineLeftSegment(),
            NewFirstLineCenterSegment(),
            NewFirstLineRightSegment(),
        },
    }
}
