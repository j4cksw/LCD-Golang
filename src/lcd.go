package lcd

func Render(number int) string {
    if number == 1 {
        return getFirstLineForNumber(1) + "\n" + getSecondLineForNumber(1) + "\n" + getThirdLineForNumber(1)
    }
    return getFirstLineForNumber(number) + "\n" + getSecondLineForNumber(number) + "\n" + getThirdLineForNumber(number)
}
