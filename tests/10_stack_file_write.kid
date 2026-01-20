// Test writing stack to file
stack output
stack output[1] = line1
stack output[2] = line2
stack output[3] = line3
file outfile = out
open outfile /tmp/kidlang_test_output.txt
write outfile stack output
close outfile
file readback = read
open readback /tmp/kidlang_test_output.txt
read readback box content
print box content
close readback
