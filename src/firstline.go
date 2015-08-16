package lcd

func NewFirstLine() Line {
    return Line{
        []Segment{
            NewFirstLineLeftSegment(),
            NewFirstLineCenterSegment(),
            NewFirstLineRightSegment(),
        },
    }
}
