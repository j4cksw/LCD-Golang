package lcd

import "fmt"

type SecondLine struct {}

func (l *SecondLine) RenderForValue(value int) string {
    return "| || |" + l.getSecondLineRightDigitForNumber(value)
}

func (l *SecondLine) getSecondLineRightDigitForNumber(number int) string {
    return fmt.Sprintf("%s%s%s",
        NewSecondLineLeftSegment().RenderForNumber(number),
        l.getSecondLineCenterSegment(number),
        l.getSecondLineRightSegment(number))
}

func (l *SecondLine) getSecondLineCenterSegment(number int) string {
    return NewSecondLineCenterSegment().RenderForNumber(number)
}

func (l *SecondLine) getSecondLineRightSegment(number int) string {
    return NewSecondLineRightSegment().RenderForNumber(number)
}
