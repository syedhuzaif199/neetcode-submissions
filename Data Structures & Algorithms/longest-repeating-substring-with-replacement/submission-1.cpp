class Solution {
public:
    int max(int *nums, int count) {
        if(count <= 0) {
            return 0;
        }
        int out = nums[0];
        for(int i = 1; i < count; i++) {
            if(nums[i] > out) {
                out = nums[i];
            }
        }
        return out;
    }
    
    int characterReplacement(string s, int k) {
        int freq[26] = {};
    
        int p = 0, q = 0;
        int freq_sum = 0;
        int longest = 0;
    
        while(q < s.size()) {
            int changes = freq_sum - max(freq, 26);
            if(changes <= k) {
                if(q-p > longest) {
                    longest = q-p;
                }
                freq[s[q] - 'A'] += 1;
                freq_sum += 1;
                q += 1;
            } else {
                freq[s[p] - 'A'] -= 1;
                freq_sum -= 1;
                p += 1;
            }
        }
    
        int changes = freq_sum - max(freq, 26);
        if(changes <= k && q-p > longest) {
            longest = q-p;
        }
    
        return longest;
    
    }
};
