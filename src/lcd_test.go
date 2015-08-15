package lcd

import "testing"

var expecteds = map[int]string {
    000: " -  -  - " + "\n" + "| || || |" + "\n" + "|_||_||_|" + "\n",
    001: " -  -    " + "\n" + "| || |  |" + "\n" + "|_||_|  |" + "\n",
    002: " -  -  - " + "\n" + "| || | _|" + "\n" + "|_||_||_ " + "\n",
}

func Test_render_000(t *testing.T) {
    for value, expected := range expecteds {
        if Render(value) != expected {
            t.Errorf("Expected \n%s but got %s", expected, Render(value))
        }
    }
}
