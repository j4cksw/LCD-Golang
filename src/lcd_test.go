package lcd

import "testing"

func Test_render_allZero_getZeroPattern(t *testing.T){
    result := Render(000)

    expectedResult := " -  -  - " + "\n" + "| || || |" + "\n" + "|_||_||_|"
    if result != expectedResult {
        t.Errorf("Expected \n%s but got %s", expectedResult, result)
    }
}
