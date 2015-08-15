package lcd

func Render(number int) string {
    firstLine := &FirstLine{}
    return firstLine.RenderForValue(number) + "\n" + getSecondLineForNumber(number) + "\n" + getThirdLineForNumber(number)
}
