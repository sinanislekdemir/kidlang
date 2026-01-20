// Test goto with multiple labels
goto second
first:
print should not print
second:
print jumped to second
goto end
third:
print should not print
end:
print reached end
