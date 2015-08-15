package lcd

func getThirdLineForNumber(number int) string {
    return "|_||_|" + getThirdLineRightDigitForNumber(number)
}

func getThirdLineRightDigitForNumber(number int) string {
    return getThirdLineLeftPiece(number) + getThirdLineCenterPiece(number) + "|"
}

func getThirdLineLeftPiece(number int) string {
    if number == 1 {
        return " "
    }
    return "|"
}

func getThirdLineCenterPiece(number int) string {
    if number == 1 {
        return " "
    }
    return "_"
}