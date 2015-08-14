package lcd

func Render(number int) string {
    if number == 1 {
        return " -  -   " + "\n" + "| || |  |" + "\n" + "|_||_|  |"
    }
    return " -  -  - " + "\n" + "| || || |" + "\n" + "|_||_||_|"
}
