package lcd

import "fmt"

type Line struct {
    leftValue string
    segments []Segment
}

func (l *Line) RenderForValue(value int) string {
    rightValue := value % 10
    value = value/10
    centerValue := value % 10
    return l.leftValue + l.getRightDigitForNumber(centerValue) + l.getRightDigitForNumber(rightValue)
}

func (l *Line) getRightDigitForNumber(number int) string {
    return fmt.Sprintf("%s%s%s",
        l.segments[0].RenderForNumber(number),
        l.segments[1].RenderForNumber(number),
        l.segments[2].RenderForNumber(number))
}
