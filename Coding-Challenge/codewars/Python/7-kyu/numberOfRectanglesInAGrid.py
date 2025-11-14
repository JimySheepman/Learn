def number_of_rectangles(m, n):
    count = 0
    for i in range(1, m+1):
        for j in range(1, n+1):
            count += i*j
    return count
#   return n*m*(n+1)*(m+1)/4
