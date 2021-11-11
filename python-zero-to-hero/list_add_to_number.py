def fun(input_list, number_to_add, count_of_number_to_add):
    if len(input_list) % 2 == 0:
        for i in [number_to_add for i in range(count_of_number_to_add)]:
            input_list.insert((len(input_list)//2), i)
        print(input_list)
    else:
        for i in [number_to_add for i in range(count_of_number_to_add)]:
            input_list.append(i)
        print(input_list)


if __name__ == '__main__':

    lst = [2, 3, 4, 5]
    lst = fun(lst, 9, 2)
    print(form(lst, 1))

    # lst = [2, 3, 4, 5, 6]
    # lst = fun(lst, 9, 2)

    # OUTPUT
    # $ python tt.py
    # [2, 3, 9, 9, 4, 5]
    # [2, 3, 4, 5, 6, 9, 9]
