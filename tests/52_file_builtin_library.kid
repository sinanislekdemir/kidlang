file source = sample
open source /tmp/kidlang_builtin_file_library.txt
print fileexists(source)
print fileread(source)
print filesize(source)
close source
