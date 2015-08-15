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

func (l *ThirdLine) getThirdLineLeftSegment(number int) string {
    if number == 1 {
        return " "
    }
    return "|"
}

func (l *ThirdLine) getThirdLineCenterSegment(number int) string {
    if number == 1 {
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
