# here is getting symptoms and solutions
def greedy_approach(a, b):
    n = len(a) - 1
    compare_list = []
    output_list = []
    # think it is algorithmically easier to start from behind. Therefore, we started it in reverse.
    while n >= 0:
        if compare_list == b:
            return output_list
        else:
            for j in range(len(a[n][2])):
                # add individual elements to a new list (tuple in array)
                compare_list.append(a[n][2][j])
            # add the tuple of the added line to a list.
            output_list.append(a[n])
            # union part
            compare_list = list(set(compare_list))
        # exit condition for loop
        n = n - 1
# finds the minimum number of solutions for medicine design
def find_min_solution(a, b):
    size = len(a)
    results = []
    id_list = []
    copy_list = a
    # call the function and replace the list itself as much as the cycle
    for m in range(size):
        copy_list = list(copy_list)
        # hold the first index
        temp = copy_list[0]
        # remove  the temp from list
        copy_list.remove(temp)
        # append the temp from list
        copy_list.append(temp)
        # call the greedy_approach function
        result = greedy_approach(copy_list, b)
        # check the result for a Null
        if result is not None:
            # append the result
            results.append(result)
            id_list.append(len(result))
    # find minimum result
    # It finds the least number of builders in the list tuple and gets the index and returns result [index]
    min_result = min(id_list)
    index = id_list.index(min_result)
    return results[index]

if __name__ == '__main__':
    symptoms_id_list = []
    solutions_id_list = []
    number_of_symptoms_that_it_fixes_list = []
    symptoms_names = ""
    full_list = []
    nested_list = []
    edit_list = []
    solutions_list = []
    best_choice_list = []
    # input the symptoms range
    loop_range_symptoms = int(input("Enter The Number Of Symptoms : "))
    # input the symptoms name and id loop & insert the list
    for i in range(0, loop_range_symptoms):
        symptom = input()
        symptoms_names = symptom.split()
        symptoms_id_list.append(int(symptoms_names[0]))
    # input the solutions range
    loop_range_solutions = int(input("Enter The Number Of Solutions : "))
    # input the solutions loop & insert the list
    for x in range(0, loop_range_solutions):
        # this line for nested array create and insert
        full_list.append([int(j) for j in input().split()])
        solutions_id_list.append(full_list[x][0])
        number_of_symptoms_that_it_fixes_list.append(full_list[x][1])
        nested_list.append(full_list[x][2:])
    # Allows us to set up tuple structure in the list
    edit_list = list(zip(solutions_id_list, number_of_symptoms_that_it_fixes_list, nested_list))
    # call the find_min_solution function
    solutions_list = find_min_solution(edit_list, symptoms_id_list)
    # takes their first index from the returned list and synchronizes them to a new list
    best_choice_list = [item[0] for item in solutions_list]
    # reverse list
    best_choice_list = best_choice_list[::-1]
    # output format layout
    print("Use solutions:", end=" ")
    if len(best_choice_list) > 1:
        for i in range(len(best_choice_list)):
            print(best_choice_list[i], end="")
            if i == len(best_choice_list) - 2:
                print(" and", end=" ")
            if i != len(best_choice_list) - 1 and i != len(best_choice_list) - 2:
                print(",", end=" ")
        print(" for your medicine.", end=" ")
    else:
        print(best_choice_list[0], end="")
        print(" for your medicine.", end=" ")