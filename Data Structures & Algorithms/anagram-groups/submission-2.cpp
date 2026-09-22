class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        auto m = std::map<std::array<int, 26>, vector<int>>();
        for(int i = 0; i < strs.size(); i++) {
            string str = strs[i];
            auto freq = getFreq(str);
            if (!m.contains(freq)) {
                m[freq] = vector<int>{};
            }
            m[freq].push_back(i);
        }

        vector<vector<string>> out{};
        for (const auto& [key, value] : m) {
            vector<string> s{};
            for(auto i : m[key]) {
                s.push_back(strs[i]);
            }
            out.push_back(s);
        }
        return out;
    }

    std::array<int, 26> getFreq(string s){
        std::array<int, 26> out{};
        for (auto c : s) {
            out[c-'a'] += 1;
        }
        return out;
    }
};