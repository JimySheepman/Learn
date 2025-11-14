def matrix_chain(p, i, j):
    if i == j:
        return 0

    m = float('inf')
    for k in range(i, j):
        count = (matrix_chain(p, i, k) + matrix_chain(p, k + 1, j) + p[i - 1] * p[k] * p[j])
        if count < m:
            m = count

    return m


if __name__ == '__main__':
    p = [6, 2, 3, 1, 3]
    n = len(p)
    i = 1
    j = n - 1

    print("Minimum number of multiplications is ", matrix_chain(p, 1, j))
