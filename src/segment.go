package lcd

const (
    SIDE_ON = "|"
    SIDE_OFF = " "
    CENTER_ON = "_"
    CENTER_OFF = " "
)

type Segment interface {
    RenderForNumber(number int) string
}

type ThirdLineLeftSegment struct {
    includes map[int]bool
}

func NewThirdLineLeftSegment() *ThirdLineLeftSegment{
    return &ThirdLineLeftSegment{
        map[int]bool {
            1: true,
            3: true,
            4: true,
            5: true,
        },
    }
}

func (s *ThirdLineLeftSegment) RenderForNumber(number int) string {
    if s.includes[number] {
        return SIDE_OFF
    }
    return SIDE_ON
}
