package lcd

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
