def double_char(s):
    double_string = ""
    for i in s:
        for j in range(2):
            double_string += i
    return double_string
# return ''.join(c * 2 for c in s)