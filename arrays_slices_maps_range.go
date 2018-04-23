// $go run arrays_slices_maps_range.go

package main

import "fmt"

func main() {   


    var a [5] int
    fmt.Println("emp:", a)



    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])


    fmt.Println("len:", len(a))
  
  ///////// OR
  
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
  
  
    var twoDim [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoDim[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoDim)
  
  }


/*
$ go run arrays.go
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2d:  [[0 1 2] [1 2 3]]

*/




/////-----------------Slicess------------------------------------ array ın büyütülebilir küçültülebilir hali gibi

    s  : =  make ([] dize ,  3 ) 
    fmt . Println ( "emp:" ,  s )      //emp: [  ]

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)      //set: [a b c]
    fmt.Println("get:", s[2])         //get: c
    fmt.Println("len:", len(s))       //len: 3

    s = append(s, "d")
    s = append(s, "e", "f")

    fmt.Println("apd:", s)         //append edilmiş hali:[a b c d e f]


    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)    // s c ye kopyalandı


     l := s[2:5]
    fmt.Println("sl1:", l)    // [c d e]   2.den  4 e kadar




    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)      //           [ g h i ]


    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)



//// OR

   var numbers []int
   printSlice(numbers)
   
   /* append allows nil slice */
   numbers = append(numbers, 0)
   printSlice(numbers)
   
   /* add one element to slice*/
   numbers = append(numbers, 1)
   printSlice(numbers)
   
   /* add more than one element at a time*/
   numbers = append(numbers, 2,3,4)
   printSlice(numbers)
   
   /* create a slice numbers1 with double the capacity of earlier slice*/
   numbers1 := make([]int, len(numbers), (cap(numbers))*2)
   
   /* copy content of numbers to numbers1 */
   copy(numbers1,numbers)
   printSlice(numbers1)   
}
   func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
       
       
/*


len = 0 cap = 0 slice = []
len = 1 cap = 2 slice = [0]
len = 2 cap = 2 slice = [0 1]
len = 5 cap = 8 slice = [0 1 2 3 4]
len = 5 cap = 16 slice = [0 1 2 3 4]

*/

////////---------------------MAPS--------------------------------
     
       
       
    m := make(map[string]int)
    m["k1"] = 7
    m["k2"] = 13
    
    fmt.Println("map:", m)   //  map: map[k1:7 k2:13]
       
    v1 := m["k1"]
    fmt.Println("v1: ", v1)      // 7
       
    fmt.Println("len:", len(m))    // 2
       
     delete(m, "k2")            // 
     
     fmt.Println("map:", m)       //    map[k1:7]
       
       
       
    n := map[string]int{"foo": 1, "bar": 2}      //map: map[foo:1 bar:2]
    fmt.Println("map:", n)  

       
///////////////--------------RANGE------------------------------------  
       
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
       fmt.Println("sum:", sum)           //sum : 9
       
       
       //or
       
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)       // index 1
        }
    }
       
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)       // a-> apple    b -> banana
    }
       
    for k := range kvs {
        fmt.Println("key:", k)          // key : a     key : b
    }
    
    for i, c := range "go" {             // 0 103         1 111
        fmt.Println(i, c)
    }
       
    
       
       
      
   
}
    



    



    







