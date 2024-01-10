class Solution(object):
    def twoSum(self, nums, target):
        for i in range (len(nums)):
            check = findValidSecond(nums[i], nums, target)
            if check != -1 and check != i:
                return [i, check]


    def findValidSecond(first, nums, target):
        for i in range (len(nums)):
            if first + nums[i] == target:
                return i
        return -1
