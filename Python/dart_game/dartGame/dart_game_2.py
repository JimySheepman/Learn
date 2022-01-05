import random
from math import sqrt

def output_function(q, t, r):  # q tuple randomly list , t sqrt list ,r tuple randomly list range
    # this function is recursive
    radius = 11
    center = (0, 0)
    if r > 0:
        print("Hit Point is:  ", q[r])
        print("Center is:  ", center)
        print("The distance is:  ", t[r])
        if t[r] > radius:
            print("Result: False")
            print("Outside of the dart board!")
        elif t[r] < radius:
            print("Result: True")
            print("Hit the board!")
            if 0 <= t[r] <= 3:
                score = 10
            elif 4 <= t[r] <= 7:
                score = 5
            elif 8 <= t[r] <= 11:
                score = 3
            elif 12 <= t[r] <= 15:
                score = 2
            elif 16 <= t[r] <= 19:
                score = 1
            else:
                score = 0
            print("Score:  ", score)
        print("------------------------------------------------")
        output_function(q, t, r - 1)  # recursive part
    else:
        return 1

def random_shot_generator(a):
    return random.randint(0, 9)

if __name__ == '__main__':
    # create a list to length 10
    my_list = list(range(10))
    # create a list to length 10
    my_list2 = list(range(10))
    # using map and create first random list
    t1 = map(random_shot_generator, my_list)
    # using map and create second random list
    t2 = map(random_shot_generator, my_list2)
    # make a randomly tuple list
    total = list(zip(t1, t2))
    # calculate to distance and return to result list
    result = list(map(sqrt, map(lambda x: x[0] ** 2 + x[1] ** 2, total)))
    print("^^** Playing Dart Game **^^")
    print("*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^")
    print("------------------------------------------------")
    # length to tuple list for sent a recursive function
    len_list = len(total)-1
    # call the output function
    output_function(total, result, len_list)

