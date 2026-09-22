class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        auto m = std::unordered_map<int, int>();

        for(int i = 0; i < nums.size(); i++) {
            int num = nums[i];
            auto it = m.find(target-num);
            if(it != m.end()) {
                return vector<int>{it->second, i};
            }
            m[num] = i;
        }
        return vector<int>{-1, -1};
    }
};