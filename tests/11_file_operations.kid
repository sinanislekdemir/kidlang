// Test basic file operations
file myfile = test
open myfile /tmp/kidlang_file_test.txt
write myfile Hello
write myfile World
close myfile
file readfile = read
open readfile /tmp/kidlang_file_test.txt
read readfile box content
print box content
close readfile
