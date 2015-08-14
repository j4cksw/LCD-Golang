package lcd

import "testing"

func Test_render_000(t *testing.T) {
    result := Render(000)

    expectedResult := " -  -  - " + "\n" + "| || || |" + "\n" + "|_||_||_|"
    if result != expectedResult {
        t.Errorf("Expected \n%s but got %s", expectedResult, result)
    }
}

func Test_render_001(t *testing.T) {
    result := Render(001)

    expectedResult := " -  -   " + "\n" + "| || |  |" + "\n" + "|_||_|  |"
    if result != expectedResult {
        t.Errorf("Expected \n%s but got \n%s", expectedResult, result)
    }
}
