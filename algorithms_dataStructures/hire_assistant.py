def hire_assistant(n):
    best = 0
    count = 0
    for i in n:
        if i > best:
            best = i
            count += 1
    return count


if __name__ == '__main__':
    a = [-1, -2, -3, -4, 6]
    result = hire_assistant(a)
    print(result)
