def insertion_sort(A):
    for j in range(1, len(A)):
        key = A[j]
        i = j - 1
        while i >= 0 and A[i] > key:
            A[i + 1] = A[i]
            i = i - 1
        A[i + 1] = key
    return A


if __name__ == '__main__':
    a = [5, 2, 4, 6, 1, 3]
    sort_a = insertion_sort(a)
    print(sort_a)
