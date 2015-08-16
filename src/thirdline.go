package lcd

import "fmt"

type ThirdLine struct {}

func (l *ThirdLine) RenderForValue(number int) string {
    return "|_||_|" + l.getThirdLineRightDigitForNumber(number)
}

func (l *ThirdLine) getThirdLineRightDigitForNumber(number int) string {
    return fmt.Sprintf("%s%s%s",
        NewThirdLineLeftSegment().RenderForNumber(number),
        NewThirdLineCenterSegment().RenderForNumber(number),
        NewThirdLineRightSegment().RenderForNumber(number))
}
