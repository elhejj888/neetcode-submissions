class Solution:
    def myPow(self, x: float, n: int) -> float:
        res = 1
        multiplier = n
        if(n<0):
            multiplier = -n
        for i in range(multiplier):
            res *= x
        if n>0:
            return res
        
        return 1/res
        