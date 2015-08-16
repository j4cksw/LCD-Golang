package lcd

import "testing"

var expecteds = map[int]string {
    0: " _  _  _ " + "\n" + "| || || |" + "\n" + "|_||_||_|" + "\n",
    1: " _  _    " + "\n" + "| || |  |" + "\n" + "|_||_|  |" + "\n",
    2: " _  _  _ " + "\n" + "| || | _|" + "\n" + "|_||_||_ " + "\n",
    3: " _  _  _ " + "\n" + "| || | _|" + "\n" + "|_||_| _|" + "\n",
    4: " _  _    " + "\n" + "| || ||_|" + "\n" + "|_||_|  |" + "\n",
    5: " _  _  _ " + "\n" + "| || ||_ " + "\n" + "|_||_| _|" + "\n",
    6: " _  _  _ " + "\n" + "| || ||_ " + "\n" + "|_||_||_|" + "\n",
    7: " _  _  _ " + "\n" + "| || |  |" + "\n" + "|_||_|  |" + "\n",
    9: " _  _  _ " + "\n" + "| || ||_|" + "\n" + "|_||_|  |" + "\n",
}

func Test_render(t *testing.T) {
    for value, expected := range expecteds {
        if Render(value) != expected {
            t.Errorf("Expected \n%sfor value %d but got\n%s", expected, value, Render(value))
        }
    }
}
