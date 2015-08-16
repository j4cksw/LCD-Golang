package lcd

type FirstLine struct {}

func NewFirstLine() Line {
    return &FirstLine{}
}

func (l *FirstLine) RenderForValue(value int) string {
    return " _  _ " + l.getFirstLineRightDigitForNumber(value)
}

func (l *FirstLine) getFirstLineRightDigitForNumber(number int) string {
    return " " + l.getFirstLineRightDigitCenterForNumber(number) + " "
}

func (l *FirstLine) getFirstLineRightDigitCenterForNumber(number int) string {
    return NewFirstLineCenterSegment().RenderForNumber(number)
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
