package lcd

type SecondLine struct {}

func (l *SecondLine) RenderForValue(value int) string {
    return "| || |" + l.getSecondLineRightDigitForNumber(value)
}

func (l *SecondLine) getSecondLineRightDigitForNumber(number int) string {
    return l.getSecondLineRightDigitLeftForNumber(number) + " |"
}

func (l *SecondLine) getSecondLineRightDigitLeftForNumber(number int) string {
    if number == 1 {
        return " "
    }
    return "|"
}

func getSecondLineForNumber(number int) string {
    return "| || |" + getSecondLineRightDigitForNumber(number)
}

func getSecondLineRightDigitForNumber(number int) string {
    return getSecondLineRightDigitLeftForNumber(number) + " |"
}

func getSecondLineRightDigitLeftForNumber(number int) string {
    if number == 1 {
        return " "
    }
    return "|"
}
