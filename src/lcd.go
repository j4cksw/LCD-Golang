package lcd

func Render(number int) string {
    firstLine := &FirstLine{}
    secondLine := &SecondLine{}
    thirdLine := &ThirdLine{}
    return firstLine.RenderForValue(number) + "\n" + secondLine.RenderForValue(number) + "\n" + thirdLine.RenderForValue(number)
}
