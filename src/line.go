package lcd

import "fmt"

type Line struct {
    leftValue string
    segments []Segment
}

func (l *Line) RenderForValue(value int) string {
    return l.leftValue + l.getRightDigitForNumber(value)
}

func (l *Line) getRightDigitForNumber(number int) string {
    return fmt.Sprintf("%s%s%s",
        l.segments[0].RenderForNumber(number),
        l.segments[1].RenderForNumber(number),
        l.segments[2].RenderForNumber(number))
}
