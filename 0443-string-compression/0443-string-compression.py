class Solution(object):
    def compress(self, chars):
        """
        :type chars: List[str]
        :rtype: int
        """
        walker, runner = 0, 0
        while runner < len(chars):
            chars[walker] = chars[runner]
            count = 1

            while runner + 1 < len(chars) and chars[runner] == chars[runner+1]:
                runner += 1
                count += 1
            
            if count > 1:
                for c in str(count):
                    chars[walker+1] = c
                    walker += 1
                
            runner += 1
            walker += 1

        return walker
        