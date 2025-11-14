def gpa_calculate(my_list):
    total_acts = 0
    sum_of_acts = 0
    total = 0
    my_l = []
    for ele in my_list:
        for j in my_list[ele]:
            total_acts += my_list[ele][j][0]
            if my_list[ele][j][1] == " ‘AA’":
                sum_of_acts += my_list[ele][j][0] * 4
            elif my_list[ele][j][1] == " ‘BA’":
                sum_of_acts += my_list[ele][j][0] * 3.5
            elif my_list[ele][j][1] == " ‘BB’":
                sum_of_acts += my_list[ele][j][0] * 3
            elif my_list[ele][j][1] == " ‘CB’":
                sum_of_acts += my_list[ele][j][0] * 2.5
            elif my_list[ele][j][1] == " ‘CC’":
                sum_of_acts += my_list[ele][j][0] * 2
            elif my_list[ele][j][1] == " ‘DC’":
                sum_of_acts += my_list[ele][j][0] * 1.5
            elif my_list[ele][j][1] == " ‘DD’":
                sum_of_acts += my_list[ele][j][0] * 1
            elif my_list[ele][j][1] == " ‘FD’":
                sum_of_acts += my_list[ele][j][0] * 0.5
            elif my_list[ele][j][1] == " ‘FF’":
                sum_of_acts += my_list[ele][j][0] * 0
            elif my_list[ele][j][1] == " ‘NA’":
                sum_of_acts += my_list[ele][j][0] * 0
        total = sum_of_acts / total_acts
        total = round(total, 2)
        print("student id: ", ele, end="        ")
        print("GPA: ", total)
        if total >= 3.00:
            my_l.append(ele)
        sum_of_acts = 0
        total_acts = 0
    return my_l

def input_information():
    continue_char = 'Y'
    student_id = 0
    course_number = 0
    full_text = ""
    all_student_information = []
    best_student = []
    student_information = {}
    course_information = {}
    while continue_char == 'Y' or continue_char == 'y':
        print("Enter student id: ", end=" ")
        student_id = input()
        student_information[student_id] = {}
        print("Enter number of course registered: ", end=" ")
        course_number = int(input())
        for element in range(0, course_number):
            print("Enter course name, ECTS, grade: ", end=" ")
            full_text = input()
            edit_text = full_text.split(",")
            edit_text[1] = int(edit_text[1])
            course_information[edit_text[0]] = tuple([edit_text[1], edit_text[2]])
        student_information[student_id].update(course_information)
        print("Do you want to continue (Y/N) ?", end=" ")
        continue_char = input()
        if continue_char == 'n' or continue_char == 'N':
            student = gpa_calculate(student_information)
            best_student.append(student)
            print("Honor Students with GPA above 3.00")
            for a in range(len(student)):
                print(student[a])
        elif continue_char == 'y' or continue_char == 'Y':
            course_information.clear()
            continue
        else:
            print("Invalid input! Try again!!!!")

if __name__ == '__main__':
    input_information()