def create_dictionary(dictionary):
    new_dict = dict()
    for i, j in dictionary.items():
        new_dict[i] = [j[0], (60*60*j[1][0])+(60*j[1][1])+(j[1][2])]
    return new_dict

def winner_runner(dictionary):
    new_dictionary = create_dictionary(dictionary)
    ordered_dictionary = sorted(new_dictionary.items(), key=lambda t: t[1][1])
    print("Gold Medal: " + ordered_dictionary[0][0][0]+' ' +
          ordered_dictionary[0][0][1]+', '+ordered_dictionary[0][1][0])
    print("Silver Medal: " + ordered_dictionary[1][0][0]+' ' +
          ordered_dictionary[1][0][1]+', '+ordered_dictionary[1][1][0])
    print("Bronze Medal: " + ordered_dictionary[2][0][0]+' ' +
          ordered_dictionary[2][0][1]+', '+ordered_dictionary[2][1][0])
    for i, j in new_dictionary.items():
        if j[0].lower() == 'turkey':
            if j[1] < 10800:
                print("Turkish athletes running under 3 hours: " +
                      i[0] + ' ' + i[1])

if __name__ == '__main__':
    marathon_dict = dict()
    while(True):
        name = tuple(map(str, input("First and last name: ").split()))
        country = input("Country: ")
        time_tuple = tuple(
            map(int, input("Finishing time in hours, minutes, seconds:").split(", ")))
        marathon_dict[name] = (country, time_tuple)
        flag = input("Do you want to continue Y/N? ")
        if flag == 'Y' or flag == 'y':
            continue
        else:
            break
    winner_runner(marathon_dict)
