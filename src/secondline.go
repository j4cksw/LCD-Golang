package lcd

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
