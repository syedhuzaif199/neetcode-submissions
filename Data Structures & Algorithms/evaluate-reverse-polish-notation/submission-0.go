func evalRPN(tokens []string) int {
    // num, err := strconv.Atoi("")

    
    stack := make([]int, len(tokens))
    top := 0

    pop2 := func () (int, int) {
        top -= 1
        a := stack[top]
        top -= 1
        b := stack[top]
        return a, b
    }

    push := func (x int) {
        stack[top] = x
        top += 1
    }

    for _, token := range tokens {
        switch token {
            case "+":
                a, b := pop2()
                push(a+b)
            case "-":
                a, b := pop2()
                push(b-a)
            case "*":
                a, b := pop2()
                push(a*b)
            case "/":
                a, b := pop2()
                push(b/a)
            default:
                num, _ := strconv.Atoi(token)
                push(num)
        }
    }

    return stack[top-1]
}

