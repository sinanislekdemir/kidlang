// Simple To-Do List
// Manage your tasks with file operations

file todos = todos.txt
open todos todo_list.txt

box count = 0

menu:
print
print ================================
print       TO-DO LIST
print ================================
print 1. Add task
print 2. View all tasks
print 3. Clear all tasks
print 4. Exit
print ================================
print Total tasks: box count
ask Choose option:

if answer = 1 then goto add_task
if answer = 2 then goto view_tasks
if answer = 3 then goto clear_tasks
if answer = 4 then goto exit

print Invalid choice!
goto menu

add_task:
print
ask Enter task description:
box task = answer
box count = box count + 1
box task_line = box count
box final = box task_line + answer
write todos box final
print ✓ Task added!
sleep 500
goto menu

view_tasks:
print
if box count = 0 then
   print No tasks yet! Add some tasks first.
   sleep 1000
   goto menu
end
print ================================
print       YOUR TASKS
print ================================
exec cat todo_list.txt
print ================================
ask Press Enter to continue...
goto menu

clear_tasks:
print
ask Are you sure you want to clear all tasks? (yes/no):

if answer != yes then
   print Cancelled.
   sleep 500
   goto menu
end

close todos
exec rm todo_list.txt
open todos todo_list.txt
box count = 0
print ✓ All tasks cleared!
sleep 500
goto menu

exit:
close todos
print
print Goodbye! Stay organized! 📝
