package lcd

import "testing"

var expecteds = map[int]string {
    000: " _  _  _ " + "\n" + "| || || |" + "\n" + "|_||_||_|" + "\n",
    001: " _  _    " + "\n" + "| || |  |" + "\n" + "|_||_|  |" + "\n",
    002: " _  _  _ " + "\n" + "| || | _|" + "\n" + "|_||_||_ " + "\n",
    003: " _  _  _ " + "\n" + "| || | _|" + "\n" + "|_||_| _|" + "\n",
    004: " _  _    " + "\n" + "| || ||_|" + "\n" + "|_||_|  |" + "\n",
    005: " _  _  _ " + "\n" + "| || ||_ " + "\n" + "|_||_| _|" + "\n",
    006: " _  _  _ " + "\n" + "| || ||_ " + "\n" + "|_||_||_|" + "\n",
}

func Test_render(t *testing.T) {
    for value, expected := range expecteds {
        if Render(value) != expected {
            t.Errorf("Expected \n%sfor value %d but got\n%s", expected, value, Render(value))
        }
    }
}
