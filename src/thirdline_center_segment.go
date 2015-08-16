package lcd

type ThirdLineCenterSegment struct {
    includes map[int]bool
}

func NewThirdLineCenterSegment() *ThirdLineCenterSegment {
    return &ThirdLineCenterSegment{
        map[int]bool {
            1:true,
            4:true,
        },
    }
}

func (s *ThirdLineCenterSegment) RenderForNumber(number int) string {
    if s.includes[number]{
        return CENTER_OFF
    }
    return CENTER_ON
}
