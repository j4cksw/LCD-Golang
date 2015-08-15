package lcd

type FirstLine struct {}

func (l *FirstLine) RenderForValue(value int) string {
    return " -  - " + l.getFirstLineRightDigitForNumber(value)
}

func (l *FirstLine) getFirstLineRightDigitForNumber(number int) string {
    return " " + l.getFirstLineRightDigitCenterForNumber(number) + " "
}

func (l *FirstLine) getFirstLineRightDigitCenterForNumber(number int) string {
    if number == 1 {
        return " "
    }
    return "-"
}

func getFirstLineForNumber(number int) string {
    return " -  - " + getFirstLineRightDigitForNumber(number)
}

func getFirstLineRightDigitForNumber(number int) string {
    return " " + getFirstLineRightDigitCenterForNumber(number) + " "
}

func getFirstLineRightDigitCenterForNumber(number int) string {
    if number == 1 {
        return " "
    }
    return "-"
}
