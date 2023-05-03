class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

class Stack:
    def __init__(self):
        self.len = 0
        self.head = None

    def peak(self):
        return self.head.data

    def pop(self):
        if self.head is None:
            self.len = 0
            return None
        self.len -= 1
        temp = self.head
        self.head = temp.next
        temp.next = None
        return temp.data

    def push(self, data):
        new_node = Node(data)
        if self.head is None:
            self.head = new_node
        else:
            new_node.next = self.head
            self.head = new_node
        self.len += 1
        

s = Stack()
s.push(1)
s.push(2)
s.push(3)
print(s.pop())
print(s.pop())
print(s.pop())
print(s.pop())
print(s.pop())
s.push(1)
print(s.pop())
print(s.pop())