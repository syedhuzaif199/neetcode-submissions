type Solution struct{}

const delim = ','

func (s *Solution) Encode(strs []string) string {
    runes := []rune{}
    for _, str := range strs {
        length := fmt.Sprint(len(str))
        for _, b := range length {
            runes = append(runes, b)
        }
        runes = append(runes, delim)
        for _, b := range str {
            runes = append(runes, b)
        }
    }
    return string(runes)
}

func (s *Solution) Decode(encoded string) []string {
    out := []string{}
    i := 0
    runes := []rune(encoded)
    for i < len(runes) {
        length := 0
        for runes[i] != delim {
            length *= 10
            length += int(runes[i] - '0')
            i+=1
        }
        i += 1
        decoded := make([]rune, 0, length)
        for length > 0 {
            decoded = append(decoded, runes[i])
            i += 1
            length -= 1
        }
        out = append(out, string(decoded))
    }

    return out
}
