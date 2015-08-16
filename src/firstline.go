package lcd

func NewFirstLine() Line {
    return Line{
        " _ ",
        []Segment{
            NewFirstLineLeftSegment(),
            NewFirstLineCenterSegment(),
            NewFirstLineRightSegment(),
        },
    }
}
