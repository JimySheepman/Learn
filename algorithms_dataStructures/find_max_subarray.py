def find_max_crossing_subarray(A, low, mid, high):
    lef_sum = float('-inf')
    sum = 0
    max_left = 0
    for i in range(mid, low - 1, -1):
        sum += A[i]
        if sum > lef_sum:
            lef_sum = sum
            max_left = i
    right_sum = float('-inf')
    sum = 0
    max_right = 0
    for j in range(mid + 1, high + 1):
        sum += A[j]
        if sum > right_sum:
            right_sum = sum
            max_right = j

    return max(max_left, max_right, lef_sum + right_sum)


def find_maximum_subarray(A, low, high):
    if high == low:
        return A[low]
    else:
        mid = (low + high) // 2
        return max(
            find_maximum_subarray(A, low, mid),
            find_maximum_subarray(A, mid + 1, high),
            find_max_crossing_subarray(A, low, mid, high)
        )


if __name__ == '__main__':
    a = [13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7]
    result = find_maximum_subarray(a, 0, len(a) - 1)
    print(result)
