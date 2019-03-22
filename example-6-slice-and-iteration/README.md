# golanglearning
golangpractices

**Slices :**

Slice is a kind of dynamic array in go which can grow and shrink.
Slice elements must of same type.

syntax:

variablename:=[] <type of slice> {comma separated values}
Eg:
movies:=[] string {"Game of Thorns ","Transporter"}

To add elements to slice use append method
Eg:
movies=append(movies,"Titanic")