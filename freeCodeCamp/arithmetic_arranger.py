def compair_text(a, b, c):
    if len(a) > len(b):
        total_len = len(a)+2
        if c == '+':
            total = int(a) + int(b)
            total = str(total)
        else:
            total = int(a) - int(b)
            total = str(total)
        for i in range(total_len-len(b)-1):
            b = ' ' + b[0:]
        for i in range(total_len-len(a)):
            a = ' ' + a[0:]
        for i in range(total_len-len(total)):
            total = ' ' + total[0:]
        ch = '-'
        t = 6*ch
        return a, c+b, t, total
    else:
        total_len = len(b)+2
        if c == '+':
            total = int(a) + int(b)
            total = str(total)
        else:
            total = int(a) - int(b)
            total = str(total)
        for i in range(total_len-len(a)):
            a = ' ' + a[0:]
        for i in range(total_len-len(b)-1):
            b = ' ' + b[0:]
        for i in range(total_len-len(total)):
            total = ' ' + total[0:]
        ch = '-'
        t = 6*ch
        return a, c+b, t, total


def print_funtion(all_list, dics):
    if not dics:
        for i in range(3):
            if i == 0:
                for j in range(len(all_list)):
                    print(all_list[j][i], end='\t')
            if i == 1:
                print('\n')
                for j in range(len(all_list)):
                    print(all_list[j][i], end='\t')
            if i == 2:
                print('\n')
                for j in range(len(all_list)):
                    print(all_list[j][i], end='\t')
    else:
        for i in range(4):
            if i == 0:
                for j in range(len(all_list)):
                    print(all_list[j][i], end='\t')
            if i == 1:
                print('\n')
                for j in range(len(all_list)):
                    print(all_list[j][i], end='\t')
            if i == 2:
                print('\n')
                for j in range(len(all_list)):
                    print(all_list[j][i], end='\t')
            if i == 3:
                print('\n')
                for j in range(len(all_list)):
                    print(all_list[j][i], end='\t')


def arithmetic_arranger(problems=list(), decison=False):
    all_list = []
    if len(problems) <= 5:
        for i in problems:
            i = i.split()
            if i[1] == '+' or i[1] == '-':
                if len(i[0]) <= 4 and len(i[2]) <= 4:
                    try:
                        all_list.append(compair_text(i[0], i[2], i[1]))
                    except:
                        print('Error: Numbers must only contain digits.')
                        return
                else:
                    print('Error: Numbers cannot be more than four digits.')
                    return
            else:
                print('Error: Operator must be '+' or '-'.')
                return
    else:
        print('Error: Too many problems.')
        return
    print_funtion(all_list, decison)


if __name__ == '__main__':

    arithmetic_arranger(["32 + 698", "3801 - 2", "45 + 43", "123 + 49"], True)
