en
// Temperature Converter
// Convert between Celsius and Fahrenheit

menu:
print
print ================================
print   TEMPERATURE CONVERTER
print ================================
print 1. Celsius to Fahrenheit
print 2. Fahrenheit to Celsius
print 3. Exit
print ================================
ask Choose conversion:

if answer = 1 then goto c_to_f
if answer = 2 then goto f_to_c
if answer = 3 then goto exit

print Invalid choice!
goto menu

c_to_f:
print
ask Enter temperature in Celsius:
box celsius = answer
box fahrenheit = box celsius * 9 / 5 + 32
print box celsius°C = box fahrenheit°F
goto menu

f_to_c:
print
ask Enter temperature in Fahrenheit:
box fahrenheit = answer
box celsius = box fahrenheit - 32 * 5 / 9
print box fahrenheit°F = box celsius°C
goto menu

exit:
print
print Thanks for using Temperature Converter!
print Goodbye!
