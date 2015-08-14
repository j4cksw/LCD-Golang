package lcd

func Render(number int) string {
    if number == 1 {
        return getFirstLineForNumber(1) + "\n" + getSecondLineForNumber(1) + "\n" + "|_||_|  |"
    }
    return getFirstLineForNumber(0) + "\n" + getSecondLineForNumber(0) + "\n" + getThirdLineForNumber(0)
}

func getThirdLineForNumber(number int) string {
    return "|_||_||_|"
}
