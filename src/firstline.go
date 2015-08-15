package lcd

type FirstLine struct {}

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
