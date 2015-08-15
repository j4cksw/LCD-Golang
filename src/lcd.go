package lcd

func Render(number int) string {
    firstLine := &FirstLine{}
    secondLine := &SecondLine{}
    return firstLine.RenderForValue(number) + "\n" + secondLine.RenderForValue(number) + "\n" + getThirdLineForNumber(number)
}
