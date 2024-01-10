class Solution(object):
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        xStr = str(x)
        if len(xStr) < 2:
            return true
        if xStr[0] == xStr[-1]:
            newStr = xStr[1:-1]
            return isPalindrome(self, newStr)
        return false
