package lcd

func Render(number int) string {
    if number == 1 {
        return getFirstLineForNumber(1) + "\n" + "| || |  |" + "\n" + "|_||_|  |"
    }
    return getFirstLineForNumber(0) + "\n" + getSecondLineForNumber(0) + "\n" + "|_||_||_|"
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

func getSecondLineForNumber(number int) string {
    return "| || || |"
}
