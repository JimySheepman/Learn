import random

def partition(A, p, r):
    x = A[r]
    i = p - 1
    for j in range(p, r, 1):
        if A[j] <= x:
            i += 1
            temp = A[i]
            A[i] = A[j]
            A[j] = temp
    temp = A[i + 1]
    A[i + 1] = A[r]
    A[r] = temp
    return i + 1


def randomized_partition(A, p, r):
    i = random.randint(p, r)
    temp = A[r]
    A[r] = A[i]
    A[i] = temp
    return partition(A, p, r)


def randomized_select(A, p, r, i):
    if p == r:
        return A[p]
    q = randomized_partition(A, p, r)
    k = q - p + 1
    if i == k:
        return A[q]
    elif i < k:
        return randomized_select(A, p, q - 1, i)
    else:
        return randomized_select(A, q + 1, r, i - k)


if __name__ == '__main__':
    a = [88, 31, 52, 25, 98, 14, 30, 62, 23, 79]
    randomized_select(a, 0, len(a) - 1, 0)
    print(a)
