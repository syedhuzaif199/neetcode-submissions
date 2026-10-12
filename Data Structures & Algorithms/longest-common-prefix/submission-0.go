func longestCommonPrefix(strs []string) string {
	i := 0
    for {
        var common rune
        for j, str := range strs {
            if i >= len(str) {
                return str[:i]
            }

            if j == 0 {
                common = []rune(str)[i]
            } else if []rune(str)[i] != common {
                return str[:i]
            }
        }
        i += 1
    }

    return strs[0] 
}
