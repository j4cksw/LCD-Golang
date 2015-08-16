package lcd

import "fmt"

type ThirdLine struct {
    segments []Segment
}

func NewThirdLine() Line {
    return &ThirdLine{
        []Segment{
            NewThirdLineLeftSegment(),
            NewThirdLineCenterSegment(),
            NewThirdLineRightSegment(),
        },
    }
}

func (l *ThirdLine) RenderForValue(number int) string {
    return "|_||_|" + l.getThirdLineRightDigitForNumber(number)
}

func (l *ThirdLine) getThirdLineRightDigitForNumber(number int) string {
    return fmt.Sprintf("%s%s%s",
        l.segments[0].RenderForNumber(number),
        l.segments[1].RenderForNumber(number),
        l.segments[2].RenderForNumber(number))
}
