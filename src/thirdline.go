package lcd

type ThirdLine struct {}

func (l *ThirdLine) RenderForValue(number int) string {
    return "|_||_|" + l.getThirdLineRightDigitForNumber(number)
}

func (l *ThirdLine) getThirdLineRightDigitForNumber(number int) string {
    return l.getThirdLineLeftPiece(number) + l.getThirdLineCenterPiece(number) + "|"
}

func (l *ThirdLine) getThirdLineLeftPiece(number int) string {
    if number == 1 {
        return " "
    }
    return "|"
}

func (l *ThirdLine) getThirdLineCenterPiece(number int) string {
    if number == 1 {
        return " "
    }
    return "_"
}
