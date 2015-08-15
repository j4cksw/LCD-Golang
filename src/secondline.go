package lcd

type SecondLine struct {}

func (l *SecondLine) RenderForValue(value int) string {
    return "| || |" + l.getSecondLineRightDigitForNumber(value)
}

func (l *SecondLine) getSecondLineRightDigitForNumber(number int) string {
    return l.getSecondLineRightDigitLeftForNumber(number) + l.getSecondLineCenterSegment(number) + "|"
}

func (l *SecondLine) getSecondLineRightDigitLeftForNumber(number int) string {
    if number == 1 || number == 2 {
        return " "
    }
    return "|"
}

func (l *SecondLine) getSecondLineCenterSegment(number int) string {
    if number == 2 {
        return "_"
    }
    return " "
}
