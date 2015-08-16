package lcd

import "fmt"

var lines = []Line{ NewFirstLine(), NewSecondLine(), NewThirdLine() }

func Render(number int) string {
    result := ""
    for _, line := range lines {
        result = fmt.Sprintf("%s%s\n", result, line.RenderForValue(number))
    }

    return result
}
