#ifndef PROCESS_H
#define PROCESS_H

#include <iostream>
#include <cstring>

#define LOW_PRIORITY 0
#define MEDIUM_PRIORITY 31
#define HIGH_PRIORITY 63

class Process {
private:
    int pid;
    char name[16];
    int priority;
    static int S_NEXT_PID;
public:
    Process(char *name = NULL, int priority = MEDIUM_PRIORITY);

    bool operator==(Process& other_process);
    bool operator<=(Process& other_process);
    bool operator<(Process& other_process);
    bool operator>(Process& other_process);

    friend std::ostream& operator<<(std::ostream& os, Process& process);
};

int Process::S_NEXT_PID = 0;

Process::Process(char *name, int priority) {
    pid = Process::S_NEXT_PID++;

    if (priority < 0) {
        this->priority = 0;
    } else if (priority > 63) {
        this->priority = 63;
    } else {
        this->priority = priority;
    }

    if (name == NULL) {
        sprintf(this->name, "Process %d", this->pid);
    }
    else if (strlen(name) > 15) {
        std::cout << "Process name length should be less than 16 bytes on Process " << this->pid << std::endl;
        sprintf(this->name, "Process %d", this->pid);
    } else {
        strncpy(this->name, name, 15);
    }
}

bool Process::operator==(Process& other_process) {
    return this->pid == other_process.pid;
}

bool Process::operator<=(Process& other_process) {
    return this->priority <= other_process.priority;
}

bool Process::operator<(Process& other_process) {
    return this->priority < other_process.priority;
}

bool Process::operator>(Process& other_process) {
    return this->priority > other_process.priority;
}

std::ostream& operator<<(std::ostream& os, Process& process) {
    os << "PROCESS NAME: " << process.name << std::endl;
    os << "PROCESS ID: " << process.pid << std::endl;
    os << "PROCESS PRIORITY: " << process.priority << std::endl;

    return os;
}


#endif
