def insertion_sort(A):
    for j in range(1, len(A)):
        key = A[j]
        i = j - 1
        while i >= 0 and A[i] > key:
            A[i + 1] = A[i]
            i = i - 1
        A[i + 1] = key
    return A


def bucket_sort(A):
    B = []
    slot = 10
    for i in range(slot):
        B.append([])

    for j in A:
        ind = int(slot * j)
        B[ind].append(j)

    for i in range(slot):
        B[i] = insertion_sort(B[i])

    k = 0
    for i in range(slot):
        for j in range(len(B[i])):
            A[k] = B[i][j]
            k += 1
    return A


if __name__ == '__main__':
    a = [0.9237, 0.45125, 0.241656, 0.211234, 0.124665, 0.357434]
    print(bucket_sort(a))
