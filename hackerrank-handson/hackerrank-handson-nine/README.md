Given a time in 12-hour AM/PM format, convert it to military (24-hour) time.

Note: Midnight is 12:00:00AM on a 12-hour clock, and 00:00:00 on a 24-hour clock. Noon is 12:00:00PM on a 12-hour clock, and 12:00:00 on a 24-hour clock.

Function Description

Complete the timeConversion function in the editor below. It should return a new string representing the input time in 24 hour format.

timeConversion has the following parameter(s):

s: a string representing time in  12-hour format


-----------------------------------------------------------------------------------

HackerLand University has the following grading policy:

Every student receives a GRADE  in the inclusive range from 0 to 100 .
Any  GRADE less than 40  is a failing grade.
Sam is a professor at the university and likes to round each student's  GRADE according to these rules:

If the  difference between the  gradeand the next multiple of  5 is less than 3, round  up to the next multiple of 5 .
If the value of  is less than 38 , no rounding occurs as the result will still be a failing grade.
For example,grade=84  will be rounded to 85  but  grade=29 will not be rounded because the rounding would result in a number that is less than 38.

Given the initial value of  grade for each of Sam's  students, write code to automate the rounding process.
--------------------------------------------------------------------------------------------------------------

Sam's house has an apple tree and an orange tree that yield an abundance of fruit. In the diagram below, the red
region denotes his house, where is s the start point, and t is the endpoint. The apple tree is to the left of his
house, and the orange tree is to its right. You can assume the trees are located on a single point, where the
apple tree is at point a , and the orange tree is at point b.
When a fruit falls from its tree, it lands d units of distance from its tree of origin along the x-axis. A negative value
of means the fruit fell d units to the tree's left, and a positive value of means it falls d units to the tree's right.
Given the value d of for m apples and n oranges, determine how many apples and oranges will fall on Sam's
house (i.e., in the inclusive range )?
