MAX = 100
i = 0
j = 0
k = 0


def matrix_multiply_recursive(row_1, column1, A, row_2, column_2, B, C):
    global i
    global j
    global k

    if i >= row_1:
        return 0

    if j < column_2:
        if k < column1:
            C[i][j] += A[i][k] * B[k][j]
            k += 1
            matrix_multiply_recursive(row_1, column1, A, row_2, column_2, B, C)

        k = 0
        j += 1
        matrix_multiply_recursive(row_1, column1, A, row_2, column_2, B, C)

    j = 0
    i += 1
    matrix_multiply_recursive(row_1, column1, A, row_2, column_2, B, C)


def matrix_multiply(row1, column1, A, row2, column2, B):
    if row2 != column1:
        print("Not Possible")
        return 0

    C = [[0 for i in range(MAX)] for i in range(MAX)]
    matrix_multiply_recursive(row1, column1, A, row2, column2, B, C)

    for i in range(row1):
        for j in range(column2):
            print(C[i][j], end=" ")
        print()


A = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
B = [[9, 8, 7], [6, 5, 4], [3, 2, 1]]

row1 = len(A[0])
col1 = len(A)
row2 = len(B[0])
col2 = len(B)

matrix_multiply(row1, col1, A, row2, col2, B)
