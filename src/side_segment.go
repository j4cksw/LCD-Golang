package lcd

const (
    SIDE_ON = "|"
    SIDE_OFF = " "
)

type SideSegment struct {
    includes map[int]bool
}

func (s *SideSegment) RenderForNumber(number int) string {
    if s.includes[number] {
        return SIDE_OFF
    }
    return SIDE_ON
}
