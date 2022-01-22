class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        counter = 0
        max_len = 0
        for num in nums:
            if num == 1:
                counter += 1
                if counter > max_len:
                    max_len = counter
            else:
                if counter > max_len:
                    max_len = counter
                counter = 0
        return max_len
        