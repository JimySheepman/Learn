def cut_rod(p, n):
    if n == 0:
        return 0
    q = float('-inf')
    for i in range(n):
        q = max(q, p[i] + cut_rod(p, n - i))

    return q


def memoized_cut_rod_aux(p, n, r):
    if r[n] >= 0:
        return r[n]
    if n == 0:
        q = 0
    else:
        q = float('-inf')
        for i in range(n):
            q = max(q, p[i] + memoized_cut_rod_aux(p, n - i, r))

    r[n] = q
    return q


def bottom_up_cut_rod(p, n):
    r = list(range(n))
    r[0] = 0
    for j in range(n):
        q = float('-inf')
        for i in range(j):
            q = max(q, p[i] + r[j - i])
        r[j] = q
    return r[n]


def extended_bottom_up_cut_rod(p, n):
    r = list(range(n))
    s = list(range(n))
    r[0] = 0
    for j in range(n-1):
        q = float('-inf')
        for i in range(j):
            if q < p[i] + r[j - i]:
                q = p[i] + r[j - i]
                s[j] = i
        r[j] = q
    return r, s


def print_cut_rod_solution(p, n):
    r, s = extended_bottom_up_cut_rod(p, n)
    while n > 0:
        print(s[n-1])
        n = n - s[n-1]


if __name__ == '__main__':
    p = [1, 5, 8, 9, 10, 17, 17, 20, 24, 30]
    n = len(p)
    print_cut_rod_solution(p, n)
