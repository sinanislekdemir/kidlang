// Test readline command
file testfile = test
open testfile /tmp/kidlang_readline_test.txt
readline testfile box line1
print box line1
readline testfile box line2
print box line2
close testfile
