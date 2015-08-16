package lcd

const SIDE_ON = "|"

type SideSegment struct {
    includes map[int]bool
}

func (s *SideSegment) RenderForNumber(number int) string {
    if s.includes[number] {
        return OFF
    }
    return SIDE_ON
}
