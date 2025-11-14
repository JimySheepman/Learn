import random
from math import sqrt


def create_shot():  # create randomly shoot function
    new_list = []
    for i in range(10):
        x = random.randint(0, 9)
        y = random.randint(0, 9)
        tup = (x, y)
        new_list.append(tup)
    return new_list


def distance_calculate(a, b):  # distance calculate function
    x1 = a[0]
    x2 = b[0]
    y1 = a[1]
    y2 = b[1]
    distance = ((x1 ** 2) - (x2 ** 2)) + ((y1 ** 2) - (y2 ** 2))
    distance = sqrt(distance)
    return distance


if __name__ == '__main__':
    center = (0, 0)
    radius = 11
    al_list = create_shot()
    print("^^** Playing Dart Game **^^")
    print("*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^")
    print("------------------------------------------------")
    print("\n")
    for x in range(len(al_list)):

        print("Hit Point is:  ", al_list[x])
        print("Center is:  ", center)
        result = distance_calculate(al_list[x], center)
        print("The distance is:  ", result)
        if result > radius:
            print("Result: False")
            print("Outside of the dart board!")
        elif result < radius:
            print("Result: True")
            print("Hit the board!")
            if 0 <= result <= 3:
                score = 10
            elif 4 <= result <= 7:
                score = 5
            elif 8 <= result <= 11:
                score = 3
            elif 12 <= result <= 15:
                score = 2
            elif 16 <= result <= 19:
                score = 1
            else:
                score = 0
            print("Score:  ", score)
        print("------------------------------------------------")
