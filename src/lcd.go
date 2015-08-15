package lcd

import "fmt"

func Render(number int) string {
    lines := []Line{ &FirstLine{}, &SecondLine{}, &ThirdLine{} }

    result := ""
    for _, line := range lines {
        result = fmt.Sprintf("%s%s\n", result, line.RenderForValue(number))
    }

    return result
}
