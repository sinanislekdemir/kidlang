EN
// Builtin helpers can work with words, files, and stacks.

box pet = rainbow robot
print contains(box pet, bow)
print replace(box pet, robot, kite)
print substring(box pet, 1, 7)

file notebook = notes
open notebook helpernotes.txt
write notebook apple|banana|carrot
print fileexists(notebook)
box note_text = fileread(notebook)
print box note_text
close notebook

stack toys
box first = stackset(toys, 1, robot)
box second = stackset(toys, 2, kite)
box key_list = stackkeys(toys)
print join(box key_list, :)
print stacklen(toys)
