package lcd

type ThirdLineRightSegment struct {
    includes map[int]bool
}

func NewThirdLineRightSegment() *ThirdLineRightSegment {
    return &ThirdLineRightSegment{
        map[int]bool{
            2: true,
        },
    }
}

func (s *ThirdLineRightSegment) RenderForNumber(number int) string {
    if s.includes[number] {
        return SIDE_OFF
    }
    return SIDE_ON
}
