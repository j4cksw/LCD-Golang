package lcd

func Render(number int) string {
    if number == 1 {
        return getFirstLineForNumber(1) + "\n" + "| || |  |" + "\n" + "|_||_|  |"
    }
    return getFirstLineForNumber(0) + "\n" + "| || || |" + "\n" + "|_||_||_|"
}

func getFirstLineForNumber(number int) string {
    // if number == 1 {
    //     return " -  -   "
    // }
    return " -  - " + getFirstLineRightDigitForNumber(number)
}

func getFirstLineRightDigitForNumber(number int) string {
    if number == 1 {
        return "   "
    }
    return " - "
}
