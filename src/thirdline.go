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

var left = map[int]bool {
    1: true,
    3: true,
    4: true,
    5: true,
}

func (l *ThirdLine) getThirdLineLeftSegment(number int) string {
    if left[number]{
        return SIDE_OFF
    }
    return SIDE_ON
}

func (l *ThirdLine) getThirdLineCenterSegment(number int) string {
    if number == 1 || number == 4{
        return " "
    }
    return "_"
}

func (l *ThirdLine) getThirdLineRightSegment(number int) string {
    if number == 2 {
        return " "
    }
    return "|"
}
