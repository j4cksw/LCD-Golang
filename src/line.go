package lcd

type Line interface {
    RenderForValue(value int) string
}
