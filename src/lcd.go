package lcd

func Render(number int) string {
    if number == 1 {
        return getFirstLineForNumber(1) + "\n" + getSecondLineForNumber(1) + "\n" + getThirdLineForNumber(1)
    }
    return getFirstLineForNumber(0) + "\n" + getSecondLineForNumber(0) + "\n" + getThirdLineForNumber(0)
}
