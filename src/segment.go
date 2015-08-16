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
