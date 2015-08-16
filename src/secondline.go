package lcd

type SecondLine struct {}

func (l *SecondLine) RenderForValue(value int) string {
    return "| || |" + l.getSecondLineRightDigitForNumber(value)
}

func (l *SecondLine) getSecondLineRightDigitForNumber(number int) string {
    return l.getSecondLineRightDigitLeftForNumber(number) + l.getSecondLineCenterSegment(number) + l.getSecondLineRightSegment(number)
}

func (l *SecondLine) getSecondLineRightDigitLeftForNumber(number int) string {
    return NewSecondLineLeftSegment().RenderForNumber(number)
}

func (l *SecondLine) getSecondLineCenterSegment(number int) string {
    return NewSecondLineCenterSegment().RenderForNumber(number)
}

func (l *SecondLine) getSecondLineRightSegment(number int) string {
    if number == 5 {
        return " "
    }
    return "|"
}
