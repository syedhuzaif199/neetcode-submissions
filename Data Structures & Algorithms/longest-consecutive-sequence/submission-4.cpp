class Solution {
public:
    struct Range {
        int start;
        int end;
    };

    int longestConsecutive(vector<int>& nums) {
        auto m = unordered_map<int, Range>();
        int longest = 0;
        for (auto num : nums) {
            if(m.contains(num)) {
                continue;
            }

            Range pre;
            Range post;
            bool pre_exists = false, post_exists = false;
            if(m.contains(num-1)) {
                pre = m[num-1];
                pre_exists = true;
            }
    
            if(m.contains(num+1)) {
                post = m[num+1];
                post_exists = true;
            }

            if(pre_exists && post_exists && num == pre.end + 1 && num == post.start - 1) {
                m.erase(pre.start);
                m.erase(pre.end);
                m.erase(post.start);
                m.erase(post.end);
                auto r = Range{
                    pre.start,
                    post.end,
                };
                m[r.start] = r;
                m[r.end] = r;
                longest = max(longest, r.end - r.start + 1);
            } else if(pre_exists && num == pre.end + 1) {
                m.erase(pre.end);
                pre.end = num;
                m[pre.start] = pre;
                m[pre.end] = pre;
                longest = max(longest, pre.end - pre.start + 1);
            } else if(post_exists && num == post.start - 1) {
                m.erase(post.start);
                post.start = num;
                m[post.start] = post;
                m[post.end] = post;
                longest = max(longest, post.end - post.start + 1);
            } else if(!post_exists && !pre_exists) {
                auto r = Range{
                    num,
                    num,
                };
                m[r.start] = r;
                m[r.end] = r;
                longest = max(longest, r.end - r.start + 1);
            }
        }
        return longest;

    }
};