n, m = map(int, input().split())
i, j, d = map(int, input().split()) # start point
graph = [list(map(int, input().split())) for _ in range(n)]

pos = [[-1, 0], [0, 1], [1, 0], [0, -1]]
cnt = 1
graph[i][j] = 2

def turn(d):
    if d == 0:
        return 3
    else:
        return d - 1

flag = True

while flag:
    for _ in range(4):
        d = turn(d)
        
        x = i + pos[d][0]
        y = j + pos[d][1]

        if 0 <= x < n and 0 <= y < m:
            if graph[x][y] == 0:
                i += pos[d][0]
                j += pos[d][1]

                graph[i][j] = 2
                cnt += 1
                break  
        else:
            continue
    
    else:
        dd = turn(d)
        dd = turn(dd)

        x = i + pos[dd][0]
        y = j + pos[dd][1]
        
        if 0 <= x < n and 0 <= y < m and graph[x][y] != 1:
            i = x
            j = y
        
        else:
             flag = False
print(cnt)
