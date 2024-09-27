# Chapter 4: Blocks, Shadows, and Control Structures

## Blocks

  - every place where a declaration occurs is a block
  
  - Variables, constants, types, and functions declared outside of any functions are placed in the
  package block. 

## Shadowing Variables

  - a shadowing variable is a variable that has the same name as a variable in a containing block

  - For as long as the shadowing variable exists, you can not access the shadowed variable

## For Statements

  - There are 4 different ways you can declare a for loop:

    ```
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
    ```

  - this ^ way is the traditional, complete for loop.

  - The Condition Only for loop (basically a while loop) 

    ```golang
        i := 1
        for i < 100 {
            fmt.Println(i)
            i = i * 2
        }
    ```

  - The Infinite For Loop

    ```golang
        for {
            fmt.Println("Hello")
        }
    ```

  - break and continue: break will end any for loop

    ```
        for {
            if !CONDITION {
                break
            }
        }
    ```

  - Continue will skip over the rest of the for loop, but not end the for loop.

  - For Range : used for iterating over elements of some of Go's builtin datatypes

    ```
        evenVals := []int{2,4,6,8,10}
        for i, v := range evenVals {
            fmt.Println(i, v) // prints the index, and the value at that index
        }
    ```

  - if you don't need the index value in that for range loop, you can use _ to catch the variable, but 
  leave it unused

## Iterating over a Map

  - maps don't have a set order when iterating over them, you won't iterate through in order from 0-n.
  
  - this is a security feature to keep people from relying on it being in order, as well as hash integrity

## Choosing the *right* for statement

  - most of the time you're going to use a for-range loop; this is the best for iterating over strings. 

  - for-range is also good for iterating slices and maps (also good with channels we will learn later)

  - use the complete for loop when you aren't iterating from the first to last element in a compound type

## Switch Statements

  - switch statements are written much like for loops, the case clauses are inside of the brackets, and 
  followed by a :

