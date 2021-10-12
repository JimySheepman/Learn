class Stack(object):

    # stack string
    def __init__(self, a):
        self.string_list = a


class PriorityQueue(object):
    def __init__(self):
        self.queue = []

    # for inserting an element in the queue
    def insert(self, data):
        self.queue.append(data)

        # for popping an element based on Priority

    def delete(self):
        try:
            min = 0
            for i in range(len(self.queue)):
                if self.queue[i] < self.queue[min]:
                    min = i
            item = self.queue[min]
            del self.queue[min]
            return item
        except IndexError:
            print()
            exit()


def heapify(arr, n, i):
    # Initialize largest as root
    largest = i
    # left = 2*i + 1
    l = 2 * i + 1
    # right = 2*i + 2
    r = 2 * i + 2
    # See if left child of root exists and is greater than root
    if l < n and arr[largest] < arr[l]:
        largest = l
    # See if right child of root exists and is greater than root
    if r < n and arr[largest] < arr[r]:
        largest = r
    # Change root, if needed
    if largest != i:
        arr[i], arr[largest] = arr[largest], arr[i]  # swap

        # Heapify the root.
        heapify(arr, n, largest)


# The main function to sort an array of given size


def heapSort(arr):
    n = len(arr)

    # Build a min-heap.
    for i in range(n // 2 - 1, -1, -1):
        heapify(arr, n, i)

    # One by one extract elements
    for i in range(n - 1, 0, -1):
        arr[i], arr[0] = arr[0], arr[i]  # swap
        heapify(arr, i, 0)


if __name__ == '__main__':

    # read form file
    f = open(r"C:\Users\Ali\Desktop\input_list.txt")
    # create an object
    myQueue = PriorityQueue()
    arry = ""
    object_list = []

    for line in f:
        # line split
        arry = line.split()
        # array of object
        object_list.append(Stack(arry))
        # object array control loop
    for i in range(len(object_list)):
        # check the request if state
        if object_list[i].string_list[0] == "I":
            # insert to event time from priority queue
            myQueue.insert(object_list[i].string_list[2])
        # check the request if state
        elif object_list[i].string_list[0] == "G":
            # delete to event time from priority queue
            myQueue.delete()
        # check the request if state
        elif object_list[i].string_list[0] == "L":
            # create a ordered list for send to heap-sort and copy priority queue
            ordered_list = myQueue.queue.copy()
            heapSort(ordered_list)
            # nested loop a search list for output
            for element in range(len(ordered_list)):
                for item in range(len(object_list)):
                    # check the string length
                    if len(object_list[item].string_list) > 1:
                        if ordered_list[element] == object_list[item].string_list[2]:
                            print(*object_list[item].string_list[1:])
                        else:
                            continue
            print("---------------------------")
        # check the request if state for error!
        else:
            print("Invalid input")
