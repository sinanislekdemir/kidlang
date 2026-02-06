## Rules!

1. Absolutely no backwards compatibility.
2. Absolutely no hacks / no workarounds / no "maybe we can do it quickly like this" / no "let's comment it out and not use it"
3. Absolutely no `git` commands to run, except `git diff`. Never roll back changes with git.
4. Change the naming of things. If a function name or variable name doesn't make sense anymore, change it. Also you can change all arguments.
5. Absolutely no code commenting for the logic. logic changes in time, comment remains, becomes misleading.
6. Never make assumptions about: import paths, model fields. Always and always check correct path and fields before writing any code.
7. All changes must pass linters.
8. Use `make` commands for building and testing.
9. Do not create any documents unless it is explicitly asked for.
10. Do not create any summary documents unless it is explicitly asked for.
11. Do not preemptively load all references - use lazy loading based on actual need.
12. Do not create any scripts to bulk-edit files. Edit them one by one if you have to.
