package lcd

type FirstLine struct {}

func NewFirstLine() Line {
    return &FirstLine{}
}

func (l *FirstLine) RenderForValue(value int) string {
    return " -  - " + l.getFirstLineRightDigitForNumber(value)
}

func (l *FirstLine) getFirstLineRightDigitForNumber(number int) string {
    return " " + l.getFirstLineRightDigitCenterForNumber(number) + " "
}

func (l *FirstLine) getFirstLineRightDigitCenterForNumber(number int) string {
    if number == 1 || number == 4{
        return " "
    }
    return "-"
}

func NewFirstLineCenterSegment() Segment{
    return &CenterSegment{
        map[int]bool{
            1: true,
            4: true,
        },
    }
}
