package lcd

const (
    CENTER_ON = "_"
    CENTER_OFF = " "
)

type Segment interface {
    RenderForNumber(number int) string
}
