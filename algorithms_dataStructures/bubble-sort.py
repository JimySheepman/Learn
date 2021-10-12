def bubble_sort(A):
    for i in range(len(A)):
        for j in range(len(A) - 1, i, -1):
            if A[j] < A[j - 1]:
                temp = A[j]
                A[j] = A[j - 1]
                A[j - 1] = temp


if __name__ == '__main__':
    a = [6, 5, 34, 2, 1, 2]
    bubble_sort(a)
    print(a)
