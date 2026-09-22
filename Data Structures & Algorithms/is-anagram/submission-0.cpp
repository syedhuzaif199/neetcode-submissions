class Solution {
public:
    bool isAnagram(string s, string t) {
        char m[256] = {};

        for(char c : s) {
            m[c] += 1;
        }

        for(char c : t) {
            m[c] -= 1;
        }

        for(int i = 0; i < 256; i++) {
            if(m[i] != 0) {
                return false;
            }
        }
        return true;
    }
};
