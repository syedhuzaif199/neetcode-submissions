class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        auto freq = unordered_map<int, int>();
		for(auto num : nums) {
			freq[num] += 1;
		}

		vector<vector<int>> bucket_list(nums.size() + 1);
		for(const auto& [key, value] : freq) {
			bucket_list[value].push_back(key);
		}

		vector<int> result;

		int current = nums.size();
		int count = k;
		while(count > 0) {
			for(auto num : bucket_list[current]) {
				if(count <= 0) {
					break;
				}
				result.push_back(num);
				count -=1;
			}
			current -= 1;
		}

		return result;
    }
};
