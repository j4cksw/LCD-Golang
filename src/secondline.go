package lcd

import "fmt"

type SecondLine struct {
    segments []Segment
}

func NewSecondLine() Line {
    return &ConcreteLine{
        "| || |",
        []Segment{
            NewSecondLineLeftSegment(),
            NewSecondLineCenterSegment(),
            NewSecondLineRightSegment(),
        },
    }
}

func (l *SecondLine) RenderForValue(value int) string {
    return "| || |" + l.getSecondLineRightDigitForNumber(value)
}

func (l *SecondLine) getSecondLineRightDigitForNumber(number int) string {
    return fmt.Sprintf("%s%s%s",
        l.segments[0].RenderForNumber(number),
        l.segments[1].RenderForNumber(number),
        l.segments[2].RenderForNumber(number))
}
