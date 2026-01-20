// Test seek command - use readline after seek
file testfile = test
open testfile /tmp/kidlang_seek_test.txt
seek testfile 2
readline testfile box line
print box line
close testfile
