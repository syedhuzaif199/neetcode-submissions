class Solution {
public:
    uint32_t reverseBits(uint32_t n) {
		int res = 0;
		int left = 32;
		while(n>0) {
			res  = res << 1;
			res = res | (n&1);
			n = n >> 1;
			left -= 1;
		}
		res = res << left;
		return res;
    }
};
