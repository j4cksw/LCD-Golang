package lcd

import "testing"

var expecteds = map[int]string {
    000: " -  -  - " + "\n" + "| || || |" + "\n" + "|_||_||_|" + "\n",
    001: " -  -    " + "\n" + "| || |  |" + "\n" + "|_||_|  |" + "\n",
    002: " -  -  - " + "\n" + "| || | _|" + "\n" + "|_||_||_ " + "\n",
    003: " -  -  - " + "\n" + "| || | _|" + "\n" + "|_||_| _|" + "\n",
    004: " -  -    " + "\n" + "| || ||_|" + "\n" + "|_||_|  |" + "\n",
}

func Test_render(t *testing.T) {
    for value, expected := range expecteds {
        if Render(value) != expected {
            t.Errorf("Expected \n%sfor value %d but got\n%s", expected, value, Render(value))
        }
    }
}
