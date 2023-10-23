import java.util.HashMap;


class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> maps = new HashMap<Integer, Integer>();
        for (int i = 0; i < nums.length; i ++) {
            if (maps.containsKey(target - nums[i])) {
                return new int[] {maps.get(target - nums[i]), i};
            } 
            maps.put(nums[i], i);
        }

        return new int[] {};
    }
}