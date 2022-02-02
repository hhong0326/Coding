def solution(diet):
    answer = 0
    answer1, answer2 = 0, 0
    idx = -1

    dp = [0]*(len(diet)+ 1)
    dp1 = [0]*(len(diet)+ 1)
    dp2 = [0]*(len(diet)+ 1)

    minn = min(diet[0])
    dp[1] = minn

    for i in range(len(diet[0])):
        if minn == diet[0][i]:
            idx = i
            break

    # mins = 1001

    for i in range(1, len(dp)):

        mins = min(diet[i-1][:idx+1])

        for j in range(3):
            if mins == diet[i-1][j]:
                idx = j
                break

        dp[i] = min(dp[i-1] + mins, mins)
        answer += dp[i]

    for i in range(1, len(dp)):

        mins = min(diet[i-1][:1])

        for j in range(3):
            if mins == diet[i-1][j]:
                idx = j
                break

        dp1[i] = min(dp1[i-1] + mins, mins)
        answer1 += dp1[i]

    for i in range(1, len(dp)):

        mins = min(diet[i-1][1:])

        for j in range(3):
            if mins == diet[i-1][j]:
                idx = j
                break

        dp2[i] = min(dp2[i-1] + mins, mins)
        answer2 += dp2[i]

    # new 
    # dp2 = [[0]*3 for _ in range(len(diet) + 1)]

    # dp2[0][0] = diet[0][0]
    # dp2[0][1] = diet[0][1]
    # dp2[0][2] = diet[0][2]
    
    # for i in range(1, len(dp2)):
    #     dp2[i][0] = min(dp2[i][0] + diet[i-1][0], diet[i-1][0])
    #     dp2[i][1] = min(dp2[i][0], dp2[i-1][1] + diet[i-1][1], diet[i-1][1]) 
    #     dp2[i][2] = min(dp2[i][0], dp2[i-1][1], dp2[i-1][2] + diet[i-1][2], diet[i-1][2])

    #     answer = min(answer, dp2[i][0], dp2[i][1], dp2[i][2])
    

    return answer