// Stack-File Integration Demo
// Shows how to read files into stacks and write stacks to files

// Read a file line-by-line into a stack
file mydata = data
open mydata /tmp/test_data.txt

read mydata stack items

print Loaded items from file:
print Item 1: stack items[1]
print Item 2: stack items[2]
print Item 3: stack items[3]

close mydata
print

// Create a stack with data
stack todos
stack todos[1] = Buy milk
stack todos[2] = Call mom
stack todos[3] = Finish homework

// Write stack to file
file output = out
open output /tmp/todolist.txt
write output stack todos
close output

print Todo list saved!
exec cat /tmp/todolist.txt
