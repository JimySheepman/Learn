def each_cons(lst, n):
    big_list = list()
    for i in range(len(lst)-n+1):
        mini_list = list()
        for j in range(n):
            mini_list.append(lst[i+j])
        big_list.append(mini_list)
    return big_list
