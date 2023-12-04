class Solution:
    def largestGoodInteger(self, num: str) -> str:
        # Initialize max_num to -1, to track the largest triplet
        max_num = -1
        
        # Set the left pointer to the start of the string
        l = 0
        
        # Set the right pointer next to the left pointer
        r = 1

        # Iterate as long as the right pointer is within the string
        while r < len(num):
            # If current characters at left and right pointers are different
            if num[l] != num[r]:
                # Move left pointer to the right pointer's position
                l = r
                
                # Increment the right pointer
                r += 1

            # If current characters at left and right pointers are the same
            else:
                # If a triplet is formed
                if r - l == 2:
                    # Update max_num with the largest digit among the found triplets
                    max_num = max(max_num, int(num[l]))
                    
                    # Move left pointer to the position after the current triplet
                    l = r + 1
                    
                    # Reset right pointer to one position after the new left pointer
                    r = l + 1
                    
                else:
                    # If a triplet is not formed yet, increment the right pointer
                    r += 1
                    
        # Return an empty string if no triplet is found, otherwise return the largest triplet
        return "" if max_num == -1 else f"{max_num}" * 3  