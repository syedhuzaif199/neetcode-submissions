type Solution struct{}

const DELIM = 'o'
const ESC   = 'b'

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	bytes := make([]byte, 0, len(strs)*len(strs[0]))
	for _, str := range strs {
		for i := range str {
			if str[i] == DELIM || str[i] == ESC {
				bytes = append(bytes, ESC)
			}
			bytes = append(bytes, str[i])
		}
		bytes = append(bytes, DELIM)
	}

	return string(bytes)
}

func (s *Solution) Decode(encoded string) []string {
	out := []string{}
	bytes := []byte{}
	esc_encountered := false
	for i := range encoded {
		c := encoded[i]
		if c == ESC && !esc_encountered {
			esc_encountered = true
			continue
		}

		if c == DELIM && !esc_encountered {
			out = append(out, string(bytes))
			bytes = []byte{}
			continue
		}

		esc_encountered = false
		bytes = append(bytes, c)
	}
	return out
}
