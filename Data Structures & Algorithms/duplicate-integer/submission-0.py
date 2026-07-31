class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        if len(nums) == 2 and nums[0] != nums[1]:
            return False
        
        elif len(nums) < 2:
            return False

        for i in range(0, len(nums)):
            for j in range(i+1, len(nums)):
                if nums[i] == nums[j]:
                    return True
        return False