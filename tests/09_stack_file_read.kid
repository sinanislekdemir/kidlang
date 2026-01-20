// Test reading file into stack
file testfile = test
open testfile /tmp/kidlang_test_data.txt
read testfile stack items
print stack items[1]
print stack items[2]
print stack items[3]
close testfile
