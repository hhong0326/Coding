n=int(input())
m=int(input())
g = [[] for _ in range(n+1)]
for _ in range(m):
    a, b = map(int,input().split())
    g[a].append(b); g[b].append(a)
    
visit = [False]*(n+1)

def dfs(s):
    stack = [s]
    
    while len(stack) != 0:
        a = stack.pop()
        visit[a] = True
        for i in g[a]:
            if not visit[i]:
                stack.append(i)
dfs(1)
print(sum(visit)-1)