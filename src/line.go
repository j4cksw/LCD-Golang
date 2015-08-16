package lcd

import "fmt"

type Line interface {
    RenderForValue(value int) string
}

type ConcreteLine struct {
    leftValue string
    segments []Segment
}

func (l *ConcreteLine) RenderForValue(value int) string {
    return l.leftValue + l.getRightDigitForNumber(value)
}

func (l *ConcreteLine) getRightDigitForNumber(number int) string {
    return fmt.Sprintf("%s%s%s",
        l.segments[0].RenderForNumber(number),
        l.segments[1].RenderForNumber(number),
        l.segments[2].RenderForNumber(number))
}
