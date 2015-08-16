package lcd

import "fmt"

type Line struct {
    leftValue string
    segments []Segment
}

func (l *Line) RenderForValue(value int) string {
    result := ""
    for i := 0; i < 3; i++ {
        digit := value%10
        value = value/10

        result = fmt.Sprintf("%s%s", l.getRightDigitForNumber(digit), result)
    }
    return result
}

func (l *Line) getRightDigitForNumber(number int) string {
    return fmt.Sprintf("%s%s%s",
        l.segments[0].RenderForNumber(number),
        l.segments[1].RenderForNumber(number),
        l.segments[2].RenderForNumber(number))
}
