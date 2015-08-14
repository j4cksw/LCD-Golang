package lcd

func Render(number int) string {
    if number == 1 {
        return getFirstLineForNumber(1) + "\n" + getSecondLineForNumber(1) + "\n" + getThirdLineForNumber(1)
    }
    return getFirstLineForNumber(0) + "\n" + getSecondLineForNumber(0) + "\n" + getThirdLineForNumber(0)
}

func getThirdLineForNumber(number int) string {
    return "|_||_|" + getThirdLineRightDigitForNumber(number)
}

func getThirdLineRightDigitForNumber(number int) string {
    if number == 1 {
        return getThirdLineLeftPiece(1) + " |"
    }
    return getThirdLineLeftPiece(0) + "_|"
}

func getThirdLineLeftPiece(number int) string {
    if number == 1 {
        return " "
    }
    return "|"
}
