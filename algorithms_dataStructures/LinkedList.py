class Node:

    def __init__(self, data):
        self.data = data  # Assign data
        self.next = None  # Initialize next as null
 
class LinkedList:
 
    def __init__(self):
        self.head = None
 
    def push(self, new_data):
 
        # 1 & 2: Allocate the Node &
        #        Put in the data
        new_node = Node(new_data)
 
        # 3. Make next of new Node as head
        new_node.next = self.head
 
        # 4. Move the head to point to new Node
        self.head = new_node
 
    def insertAfter(self, prev_node, new_data):
 
        # 1. check if the given prev_node exists
        if prev_node is None:
            print ("The given previous node must inLinkedList.")
            
 
        #  2. create new node &
        #      Put in the data
        new_node = Node(new_data)
 
        # 4. Make next of new Node as next of prev_node
        new_node.next = prev_node.next
 
        # 5. make next of prev_node as new_node
        prev_node.next = new_node
 
    def append(self, new_data):
 
        # 1. Create a new node
        # 2. Put in the data
        # 3. Set next as None
        new_node = Node(new_data)
 
        # 4. If the Linked List is empty, then make the
        #    new node as head
        if self.head is None:
            self.head = new_node
            return
 
        # 5. Else traverse till the last node
        last = self.head
        while (last.next):
            last = last.next
 
        # 6. Change the next of last node
        last.next =  new_node
 
    def printList(self):
        temp = self.head
        while (temp):
            print (temp.data)
            temp = temp.next

    def push(self, new_data):
        new_node = Node(new_data)
        new_node.next = self.head
        self.head = new_node
    
    def deleteList(self):
    
            # initialize the current node
            current = self.head
            while current:
                prev = current.next  # move next node
    
                # delete the current node
                del current.data
    
                # set current equals prev node
                current = prev
    
            # In python garbage collection happens
            # therefore, only
            # self.head = None

    def deleteNode(self, key):
         
        # Store head node
        temp = self.head
 
        # If head node itself holds the key to be deleted
        if (temp is not None):
            if (temp.data == key):
                self.head = temp.next
                temp = None
                return
 
        # Search for the key to be deleted, keep track of the
        # previous node as we need to change 'prev.next'
        while(temp is not None):
            if temp.data == key:
                break
            prev = temp
            temp = temp.next
 
        # if key was not present in linked list
        if(temp == None):
            return
 
        # Unlink the node from linked list
        prev.next = temp.next
 
        temp = None

    def getCountRec(self, node):
        if (not node): # Base case
            return 0
        else:
            return 1 + self.getCountRec(node.next)
 
    def getCount(self):
       return self.getCountRec(self.head)
    
    def search(self, x):
 
        # Initialize current to head
        current = self.head
 
        # loop till current not equal to None
        while current != None:
            if current.data == x:
                return True # data found
             
            current = current.next
         
        return False # Data Not found
    
        # Returns data at given index in linked list

    def getNth(self, index):
        current = self.head  # Initialise temp
        count = 0  # Index of current node
 
        # Loop while end of linked list is not reached
        while (current):
            if (count == index):
                return current.data
            count += 1
            current = current.next
 
        # if we get to this line, the caller was asking
        # for a non-existent element so we assert fail
        assert(False)
        return 0

    def count(self, search_for):
        current = self.head
        count = 0
        while(current is not None):
            if current.data == search_for:
                count += 1
            current = current.next
        return count

    def detectLoop(self):
        s = set()
        temp = self.head
        while (temp):
 
            # If we have already has
            # this node in hashmap it
            # means their is a cycle
            # (Because you we encountering
            # the node second time).
            if (temp in s):
                return True
 
            # If we are seeing the node for
            # the first time, insert it in hash
            s.add(temp)
 
            temp = temp.next
 
        return False

    def removeDuplicates(self):
        temp = self.head
        if temp is None:
            return
        while temp.next is not None:
            if temp.data == temp.next.data:
                new = temp.next.next
                temp.next = None
                temp.next = new
            else:
                temp = temp.next
        return self.head

    def swapNodes(self, x, y):
 
        # Nothing to do if x and y are same
        if x == y:
            return
 
        # Search for x (keep track of prevX and CurrX)
        prevX = None
        currX = self.head
        while currX != None and currX.data != x:
            prevX = currX
            currX = currX.next
 
        # Search for y (keep track of prevY and currY)
        prevY = None
        currY = self.head
        while currY != None and currY.data != y:
            prevY = currY
            currY = currY.next
 
        # If either x or y is not present, nothing to do
        if currX == None or currY == None:
            return
        # If x is not head of linked list
        if prevX != None:
            prevX.next = currY
        else:  # make y the new head
            self.head = currY
 
        # If y is not head of linked list
        if prevY != None:
            prevY.next = currX
        else:  # make x the new head
            self.head = currX
 
        # Swap next pointers
        temp = currX.next
        currX.next = currY.next
        currY.next = temp

    def moveToFront(self):
        tmp = self.head
        sec_last = None # To maintain the track of
                        # the second last node
  
        # To check whether we have not received 
        # the empty list or list with a single node
        if not tmp or not tmp.next: 
            return
  
        # Iterate till the end to get
        # the last and second last node 
        while tmp and tmp.next :
            sec_last = tmp
            tmp = tmp.next
  
        # point the next of the second
        # last node to None
        sec_last.next = None
  
        # Make the last node as the first Node
        tmp.next = self.head
        self.head = tmp

    def reverse(self):
        prev = None
        current = self.head
        while(current is not None):
            next = current.next
            current.next = prev
            prev = current
            current = next
        self.head = prev
 
