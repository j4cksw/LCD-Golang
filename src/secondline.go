package lcd

import "fmt"

type SecondLine struct {}

func (l *SecondLine) RenderForValue(value int) string {
    return "| || |" + l.getSecondLineRightDigitForNumber(value)
}

func (l *SecondLine) getSecondLineRightDigitForNumber(number int) string {
    return fmt.Sprintf("%s%s%s",
        NewSecondLineLeftSegment().RenderForNumber(number),
        NewSecondLineCenterSegment().RenderForNumber(number),
        NewSecondLineRightSegment().RenderForNumber(number))
}
