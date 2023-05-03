class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

class Queue:
    def __init__(self):
        self.len = 0
        self.head = None
        self.tail = None

    def peak(self):
        return self.head.data

    def dequeue(self):
        if self.head is None:
            self.tail = None
            self.len = 0
            return None
        self.len -= 1
        temp = self.head
        self.head = self.head.next
        temp.next = None
        return temp.data
    
    def enqueue(self, data):
        new_node = Node(data)
        if self.tail is None:
            self.head = self.tail = new_node
        else:
            self.tail.next = new_node
            self.tail = new_node
        self.len += 1
        

q = Queue()
q.enqueue(1)
q.enqueue(2)
q.enqueue(3)
print(q.dequeue())
print(q.dequeue())
print(q.dequeue())
print(q.dequeue())
print(q.dequeue())