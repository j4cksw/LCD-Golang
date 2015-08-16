package lcd

type CenterSegment struct {
    includes map[int]bool
}

func (s *CenterSegment) RenderForNumber(number int) string {
    if s.includes[number]{
        return CENTER_OFF
    }
    return CENTER_ON
}
