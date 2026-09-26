class Solution {
    static class Range {
        int start;
        int end;

        Range(int start, int end) {
            this.start = start;
            this.end = end;
        }
    }

    public int longestConsecutive(int[] nums) {
        Map<Integer, Range> m = new HashMap<>();
        int longest = 0;

        for (int num : nums) {
            Range r = m.get(num);

            if (r != null) {
                continue;
            }

            Range pre = m.get(num - 1);
            Range post = m.get(num + 1);

            if (pre != null
                    && post != null
                    && num == pre.end + 1
                    && num == post.start - 1) {

                m.remove(pre.start);
                m.remove(pre.end);
                m.remove(post.start);
                m.remove(post.end);

                r = new Range(pre.start, post.end);

                m.put(r.start, r);
                m.put(r.end, r);

                longest = Math.max(longest, r.end - r.start + 1);

            } else if (pre != null && num == pre.end + 1) {

                m.remove(pre.end);

                pre.end = num;

                m.put(pre.start, pre);
                m.put(pre.end, pre);

                longest = Math.max(
                    longest,
                    pre.end - pre.start + 1
                );

            } else if (post != null && num == post.start - 1) {

                m.remove(post.start);

                post.start = num;

                m.put(post.start, post);
                m.put(post.end, post);

                longest = Math.max(
                    longest,
                    post.end - post.start + 1
                );

            } else if (post == null && pre == null) {

                r = new Range(num, num);

                m.put(r.start, r);
                m.put(r.end, r);

                longest = Math.max(
                    longest,
                    r.end - r.start + 1
                );
            }
        }

        return longest;
    }
}