class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1 or numRows >= len(s):
            return s
        
        rows = [""] * numRows
        count = 0
        direction = 1

        for c in s:
            rows[count] += c
            count += direction
            if count == numRows - 1 or count == 0:
                direction *= -1
        
        return "".join(rows)
