package lcd

func NewFirstLine() Line {
    return Line{
        " _  _ ",
        []Segment{
            NewFirstLineLeftSegment(),
            NewFirstLineCenterSegment(),
            NewFirstLineRightSegment(),
        },
    }
}
