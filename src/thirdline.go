package lcd

import "fmt"

type ThirdLine struct {}

func (l *ThirdLine) RenderForValue(number int) string {
    return "|_||_|" + l.getThirdLineRightDigitForNumber(number)
}

func (l *ThirdLine) getThirdLineRightDigitForNumber(number int) string {
    return fmt.Sprintf("%s%s%s",
        l.getThirdLineLeftSegment(number),
        l.getThirdLineCenterSegment(number),
        l.getThirdLineRightSegment(number))
}

func (l *ThirdLine) getThirdLineLeftSegment(number int) string {
    return NewThirdLineLeftSegment().RenderForNumber(number)
}

func (l *ThirdLine) getThirdLineCenterSegment(number int) string {
    return NewThirdLineCenterSegment().RenderForNumber(number)
}

func (l *ThirdLine) getThirdLineRightSegment(number int) string {
    return NewThirdLineRightSegment().RenderForNumber(number)
}
