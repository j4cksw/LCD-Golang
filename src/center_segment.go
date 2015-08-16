package lcd

const CENTER_ON = "_"

type CenterSegment struct {
    includes map[int]bool
}

func (s *CenterSegment) RenderForNumber(number int) string {
    if s.includes[number]{
        return OFF
    }
    return CENTER_ON
}
