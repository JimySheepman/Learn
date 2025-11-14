def minimum(A):
    _min = A[0]
    for i in range(1, len(A)):
        if _min > A[i]:
            _min = A[i]
    return _min


def maximum(A):
    _max = A[0]
    for i in range(1, len(A)):
        if _max < A[i]:
            _max = A[i]
    return _max


if __name__ == '__main__':
    a = [13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7]
    print(maximum(a))
    print(minimum(a))
