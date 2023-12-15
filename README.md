# aqaryTest_Scripts

**Question # 2**

Given a string s, rearrange the characters of s so that any two adjacent characters are not the same.
Return any possible rearrangement of s or return "" if not possible.

**Solution:**
```
    file           : string_arrange.go
    how to run     : go run string_arrange.go
```

**Question# 3**

Table: Seat (id int, student varchar)

id is the primary key (unique value) column for this table.
Each row of this table indicates the name and the ID of a student.
id is a continuous increment.

Write a solution to swap the seat id of every two consecutive students. If the number of students is odd,
the id of the last student is not swapped.
Return the result table ordered by id in ascending order.

**Solution:** 
```
    file           : seat_shuffle/seat_shuffle.go
    how to run     : cd seat_shuffle && go run seat_shuffle.go
    requirement    : One postgres instance running on localhost with port 5432
```

**Question# 4**
Assume there are M goroutines reading from a shared buffer (such as a byte slice) and N goroutines
writing into it. How to avoid deadlock and race condition? The goroutines are always running and there
is no wait group. Using wait groups is not allowed.
Solve for M = 8 and N = 2
Solve for M = 8 and N = 8
Solve for M = 8 and N = 16
Solve for M = 2 and N = 8

**Solution:**
```
    file : goroutine_racecondition.go
    how to run: go run goroutine_racecondition
```
