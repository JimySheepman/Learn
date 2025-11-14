def get_score(n):
    result = 0
    for i in range(1, n+1):
        result += i*50
    return result