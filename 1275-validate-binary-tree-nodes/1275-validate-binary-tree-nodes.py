import collections

class Solution(object):
    def validateBinaryTreeNodes(self, n, leftChild, rightChild):
        """
        :type n: int
        :type leftChild: List[int]
        :type rightChild: List[int]
        :rtype: bool
        """
        parent_to_children = collections.defaultdict(list)
        child_to_parent = {}

        for node in range(n):
            left = leftChild[node]
            right = rightChild[node]

            if left != -1:
                parent_to_children[node].append(left)

                if left not in child_to_parent:
                    child_to_parent[left] = node
                else:
                    return False
                    
            if right != -1:
                parent_to_children[node].append(right)

                if right not in child_to_parent:
                    child_to_parent[right] = node
                else:
                    return False
        
        root_candidate = None
        for node in range(n):
            if node not in child_to_parent:
                if not root_candidate:
                    root_candidate = node
                else:
                    return False
        
        if root_candidate is None:
            return False

        visited = set()
        queue = collections.deque([root_candidate])

        while queue:
            cur_node = queue.popleft()

            visited.add(cur_node)
            for child in parent_to_children[cur_node]:
                queue.append(child)
        
        return len(visited) == n