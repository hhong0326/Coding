N = int(input())
maze = [list(map(str, input())) for _ in range(N)]
    
pos = [[-1, 0], [0, 1], [1, 0], [0, -1]]

def bfs(q):
    cnt = 0
    while(len(q) != 0):
        
        a = q.pop(0)
        if a[0] < 0 or a[1] < 0 or a[0] >= N or a[1] >= N or maze[a[0]][a[1]] != "1":
            continue
        
        maze[a[0]][a[1]] = "2"
        cnt += 1
        
        for k in range(len(pos)):
            q.append([a[0] + pos[k][0], a[1] + pos[k][1]])
    return cnt
    
arr = []
for i in range(N):
    for j in range(N):
        if maze[i][j] == "1":
            count = bfs([[i, j]])
            arr.append(count)

arr.sort()
for i in range(len(arr)):
    print(str(arr[i]), end = " ")