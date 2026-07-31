class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        strList = ""
        for i in range(len(digits)):
            strList+=str(digits[i])
        
        res = int(strList)+1
        strList = list(str(res))
        return [int(item) for item in strList]
        