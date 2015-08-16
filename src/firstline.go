package lcd

type FirstLine struct {}

func NewFirstLine() Line {
    return &FirstLine{}
}

func (l *FirstLine) RenderForValue(value int) string {
    return " _  _ " + l.getFirstLineRightDigitForNumber(value)
}

func (l *FirstLine) getFirstLineRightDigitForNumber(number int) string {
    return NewFirstLineLeftSegment().RenderForNumber(number) + NewFirstLineCenterSegment().RenderForNumber(number) + NewFirstLineRightSegment().RenderForNumber(number)
}

func NewFirstLineCenterSegment() Segment{
    return &CenterSegment{
        map[int]bool{
            1: true,
            4: true,
        },
    }
}

func NewFirstLineLeftSegment() Segment{
    return &AlwaysOffSegment{}
}

func NewFirstLineRightSegment() Segment{
    return &AlwaysOffSegment{}
}
