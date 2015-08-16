package lcd

import "fmt"

type Line struct {
    leftValue string
    segments []Segment
}

const MAX_DIGIT = 3

func (l *Line) RenderForValue(value int) string {
    var result string
    for i := 0; i < MAX_DIGIT; i++ {
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
