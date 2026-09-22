type Solution struct{}

const DELIMITER = 'W'
const ESCAPE = 'o'

func (s *Solution) Encode(strs []string) string {
    lens := func(strs []string) int {
        out := 0
        for _, str := range strs {
            out += len(str)
        }
        return out
    }
    out := make([]byte, 0, lens(strs) + len(strs))
    for _, str := range strs {
        for i := 0; i < len(str); i += 1 {
            if str[i] == DELIMITER || str[i] == ESCAPE {
                out = append(out, ESCAPE)
            }
            out = append(out, str[i])
        }
        out = append(out, DELIMITER)
    }
    return string(out)
}

func (s *Solution) Decode(encoded string) []string {
    out := []string{}
    bytes := []byte{}
    encountered_esc := false
    for _, c := range []byte(encoded) {
        if c == ESCAPE {
            if !encountered_esc {
                encountered_esc = true
                continue
            }
        }
        if c == DELIMITER {
            if !encountered_esc {
                out = append(out, string(bytes))
                bytes = []byte{}
                continue
            }
        }

        bytes = append(bytes, c)
        encountered_esc = false
    }
    if len(bytes) > 0 {
        out = append(out, string(bytes))
    }
    return out
}
