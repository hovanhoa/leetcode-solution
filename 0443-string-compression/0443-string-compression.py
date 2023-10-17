class Solution:
    def compress(self, chars: List[str]) -> int:
        result, cnt = 0, 1
        j = 0
        for i in range(1, len(chars)+1): 
            if i < len(chars) and chars[i] == chars[i-1]:
                cnt += 1
            else:
                chars[j] = chars[i-1]
                j += 1

                if cnt == 1:
                    continue
                for ch in str(cnt):
                    chars[j] = ch
                    j += 1

                cnt = 1
        return j  
