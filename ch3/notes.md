## Composite Types

### Declaring a slice with default values:
    `data := []int{2,4,6,8}`

### Three posibilities for declaring length with a slice:

    1. if you are using a slice as a buffer, then specify a non-zero length
    2. if you are *sure* you know the exact size you want, you can specify the length and index
       into the slice to set values
    3. in other situations, use `make` with a zero length and a specified capacity.

### Slicing Slices

    - a slice expression creates a slice from a slice. Written with: `[starting offset:end offset]`
    - starting offset is the first index that is included, end offset is one past the last position to use
    - when you take a slice from a slice, you are not making a copy of the data. You just have
    two variables sharing memory. This means that changes to an element in a slice affect all slices that
    share that element. 
    - to avoid complicated slice situations, you should either never use `append` with a subslice or make
    sure that append doesn't cause an override by using a full slice expression. 
    - a full slice expression looks like this: `slice_var[start offset: end offset: last position]`

        - this `last position` is the last position in the parent slices capacity that is available
        to the new subslice. 

    - to create a slice independent of its parent, use the `copy` function. 
        `num := copy(destination slice, source slice)`

### Maps

    - maps are useful when you want to associate one value to another
    - they are written in the form of `map[keyType]valueType` 
    - declarations:

        `var nilMap map[string]int` : this creates a nil map, attempting to read it returns 0 while
        attempting to write to it causes a panic

        `totalWins := map[string]int` : this is not the same as a nil map. You can write and read to this one.

        ```
            teams := map[string][]string {
                "Orcas": []string{"Fred, "Ralph", "Bijou"},
                "Lions": []string{"Sarah", "Peter", "Billie"}
            }
        ```

        `ages := make(map[int][]string, 10)` : this is with make, it has an initial length of 0
        but can grow past the initial capacity

    - a map is a data structure that associates one value to another. Maps built in Go are *hash maps* 
    - when reading from a map, you can use a *comma ok* idiom to read the the value of the key to the first
    variable, and if there is no key value the second variable gets a `false`, if it is in there: it gets 
    `true`

        ```
            m := make(map[string]int{
                "hello": 5,
                "world": 0,
            }

            v, ok := m["hello"] // v = 5, ok = true
            v, ok := m["goodbye"] // v = 0, ok = false
        ```

### Structs

    - when you have related data you want to store together, use a struct
    - simple declaration:

        ```
            type person struct {
                name string
                age int
                pet string
            }
        ``` 
    
    - initializing a new struct literal:
    
        ```
            julia := person{
                "Julia", // or name: "Julia",
                40, // or age: 40,
                "cat", // or pet: "cat"
            }
        ```

    - fields in an initialized struct are accessed with dot notation: `fmt.Println(julia.age) // prints 40`

    - 
