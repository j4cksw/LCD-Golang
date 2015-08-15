package lcd

import "testing"

func Test_render_000(t *testing.T) {
    result := Render(000)

    expectedResult := " -  -  - " + "\n" + "| || || |" + "\n" + "|_||_||_|" + "\n"
    if result != expectedResult {
        t.Errorf("Expected \n%s but got %s", expectedResult, result)
    }
}

func Test_render_001(t *testing.T) {
    result := Render(001)

    expectedResult := " -  -    " + "\n" + "| || |  |" + "\n" + "|_||_|  |" + "\n"
    if result != expectedResult {
        t.Errorf("Expected \n%s but got \n%s", expectedResult, result)
    }
}

func Test_render_002(t *testing.T) {
    result := Render(002)

    expectedResult := " -  -  - " + "\n" + "| || | _|" + "\n" + "|_||_||_ " + "\n"
    if result != expectedResult {
        t.Errorf("Expected \n%s but got \n%s", expectedResult, result)
    }
}
