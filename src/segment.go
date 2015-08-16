package lcd

type Segment interface {
    RenderForNumber(number int) string
}
