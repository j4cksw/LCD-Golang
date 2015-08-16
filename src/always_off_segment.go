package lcd

const OFF = " "

type AlwaysOffSegment struct {}

func (s *AlwaysOffSegment) RenderForNumber(number int) string {
    return OFF
}
