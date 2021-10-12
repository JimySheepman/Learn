#include "BinarySearchTree.hpp"
#include "Student.hpp"
#include <iostream>
#include <fstream>
#include <vector>
#include <boost/algorithm/string.hpp>

int main() {
    BinarySearchTree<Student> bst_student;
    std::ifstream csv_file;
    csv_file.open("hw3_text.csv", std::ios::in);

    std::string line;
    while (std::getline(csv_file, line)) {
        std::vector<std::string> result;
        boost::split(result, line, boost::is_any_of(","));

        Student *s = new Student;
        s->id = std::stoi(result[0]);
        s->name = result[1];
        s->surname = result[2];
        s->grade = std::stoi(result[3]);

        bst_student.Insert(*s);
    }

    csv_file.close();

    bst_student.PostorderTraversal();
    std::cout << std::endl;

    return 0;
}