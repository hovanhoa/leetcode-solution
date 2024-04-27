from queue import Queue

class Solution:
    def nearestExit(self, maze: List[List[str]], entrance: List[int]) -> int:
        
        q = Queue()
        q.put((entrance[0] , entrance[1] , 0))
        t = 0
        ans =[]
        maze[entrance[0]][entrance[1]] = "+"
 
        while(not q.empty()) :
            s = q.qsize()
            for i in range(s) :
                node = q.get() 
                # boundary conditions
                if node[0]==0 or node[0]==len(maze)-1 :
                    if node[2]!=0 :
                        return node[2]
                    ans += [node[2]]
                elif node[1] == 0 or node[1]==len(maze[0])-1 :
                    if node[2]!=0 :
                        return node[2]
                    ans += [node[2]]
                
                dx = [0 , -1 , 0 ,1 ]
                dy = [-1 , 0 , 1 , 0]

                for ind in range(4) :
                    x = node[0] + dx[ind]
                    y = node[1] + dy[ind] 
                    if x>=0 and x<len(maze) and y>=0 and y<len(maze[0]) and maze[x][y]!="+" :
                        maze[x][y] = "+"
                        q.put((x,y,node[2]+1))

        if len(ans) == 0:
            return -1 
    
        return -1 if max(ans)==0 else max(ans) 