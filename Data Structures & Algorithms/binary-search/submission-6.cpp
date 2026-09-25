class Solution {
public:
    int search(vector<int>& nums, int target) {
        int p = 0, q = nums.size() - 1;
        while(p <= q) {
            int mid = p + (q-p)/2;
            if (nums[mid] == target) {
                return mid;
            }

            if (nums[mid] < target) {
                p = mid + 1;
            } else {
                q = mid - 1;
            }
        }
        return -1;
    }
};
