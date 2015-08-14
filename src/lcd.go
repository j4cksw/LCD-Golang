package lcd

func Render(number int) string {
    return getFirstLineForNumber(number) + "\n" + getSecondLineForNumber(number) + "\n" + getThirdLineForNumber(number)
}
