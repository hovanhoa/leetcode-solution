class Solution:
    def sortVowels(self, s: str) -> str:
        # Step 1: Collect vowels and sort them in descending order
        vowels_sorted = sorted([c for c in s if c.lower() in 'aeiou'], reverse=True)

        # Step 2: Construct the answer string by replacing vowels in sorted order
        result = []
        for char in s:
            if char.lower() in 'aeiou':
                result.append(vowels_sorted.pop())
            else:
                result.append(char)

        # Step 3: Join the characters to form the final string
        return ''.join(result)