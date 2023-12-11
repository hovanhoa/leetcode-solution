class Solution:
    def findSpecialInteger(self, arr: List[int]) -> int:
        def binary_search(target, is_first):
            left, right = 0, len(arr) - 1
            result = -1

            while left <= right:
                mid = left + (right - left) // 2

                if arr[mid] == target:
                    result = mid
                    if is_first:
                        right = mid - 1
                    else:
                        left = mid + 1
                elif arr[mid] < target:
                    left = mid + 1
                else:
                    right = mid - 1

            return result
        
        n = len(arr)
        quarter = n // 4

        # Handle the case where quarter is zero
        if quarter == 0:
            return arr[0] if n > 0 else None

        # Check every possible candidate element
        for i in range(quarter, n, quarter):
            # Use binary search to find the first and last occurrence of the candidate element
            left_occurrence = binary_search(arr[i], True)
            right_occurrence = binary_search(arr[i], False)

            # Check if the frequency is greater than 25%
            if right_occurrence - left_occurrence + 1 > quarter:
                return arr[i]