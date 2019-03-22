# golanglearning
golangpractices

go method syntax

<func keyword> <functionName>() <returntype> {
    return <returnvalue>
}

Example : 
func getMovieName() string{
    return "ageofthorns"


**Note :** If a function returns a value we always need to mark the type that function is returning
else not require to specify the return type.

Files in same package can freely call methods defined in one another
This example demonstrate this point as well.

To run these files use command go run file1.go file2.go
