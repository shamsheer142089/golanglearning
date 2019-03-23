# golanglearning
golangpractices

#### **Go Routene**


Go routine is an independently running function initiated 
by go .

In general large programs are classified into small sub programs.

Each sub-program might be a thread or function.

Please note that go routine is a small function but not a thread.
It has its own stack which is cheap in memory and it either grow or shrink

Go routine is for concurrent execution . Go routine communicate each other through 
channels.

main function is always an implicit routine in go.

When you add "go" key word before a function thats routine in go.

This examplse demonstrated in this repo will explain the concepts in the below order

1. Basic Routine Concept
2. How routines communicate 


**1. Basic Routine Concept**

The example in the file **example-goroutine-handson-one.go**
will loop through 10 numbers and find whether number is even or odd using go.

The program compiled and run was successful.
Its displaying absurd results for every run as given below.

**$ go run example-goroutene-handson-one.go

Iam Odd :  1

Iam Even :  0

Iam Even :  2

Iam Even :  4


$ go run example-goroutene-handson-one.go

Iam Even :  2**

The problem is we dont have proper communication among routine calls.
Here comes the channel to rescue which is demonstrated in the next example.

2. How routines communicate
 We need to create a channel to make communication among the various routines.
 
        a.) Create a channel
 
        b.) Pass the channel to thego routine
 
        c.) Handle the response from channel in man routine.
 
 This program handles all the above mentioned ..
 
 
 


