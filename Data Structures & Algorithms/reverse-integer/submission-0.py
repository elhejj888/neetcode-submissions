class Solution:
    def reverse(self, x: int) -> int:
        if x < -2**31 or x > (2**31)-1:
            return 0

        x_str = str(x) if x>0 else str(x * -1)
        res = ""
        for i in range(len(x_str) - 1,-1,-1 ):
            res+=x_str[i]

        res = int(res) if x > 0 else int(res) * -1

        return 0 if (res <= -2 ** 31 or res >= 2**31 -1) else res

        