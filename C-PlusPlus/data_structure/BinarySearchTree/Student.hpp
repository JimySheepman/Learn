#ifndef STUDENT_H
#define STUDENT_H

#include <string>
#include <iostream>

class Student {
public:
    int id;
    std::string name;
    std::string surname;
    int grade;
    bool operator>(const Student& other) {
        return grade > other.grade;
    }

    bool operator==(const Student& other) {
        return other.grade == grade;
    }

    friend std::ostream& operator<<(std::ostream& os, const Student& std) {
        os << std.id << ": " << std.name + " " + std.surname << " - " << std.grade << " --> " << std::endl;
        return os;
    }
};

#endif