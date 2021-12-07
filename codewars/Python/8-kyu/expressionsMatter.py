def expression_matter(a, b, c):
    result = 0
    count = 0
    result = a*(b+c)
    if result > count:
        count = result
    result = a*b*c
    if result > count:
        count = result
    result = a+b+c
    if result > count:
        count = result
    result = a+b*c
    if result > count:
        count = result
    result = (a+b) * c
    if result > count:
        count = result
    return count
