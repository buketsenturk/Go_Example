//you run ====>>>> $go run variables_const_for_if-else_switch.go

package main

import "fmt"
import "time"

func main() {
    
    var a = "string"
    fmt.Println(a)
    
    
   var x float64
   x = 15.6
   fmt.Println(x)
   fmt.Printf("x is of type %T\n", x)
   
   
    var d = true
    fmt.Println(d)
    
    var e int
    fmt.Println(e)       // 0 
    
    
    f := "string2"
    fmt.Println(f)

    
    
    
    var k, l, m = 3, 4, "foo"  
	
   fmt.Println(k)
   fmt.Println(l)
   fmt.Println(m)
   fmt.Printf("a is of type %T\n", k)       //k is of type int
   fmt.Printf("b is of type %T\n", l)       //l is of type int
   fmt.Printf("c is of type %T\n", m)       //m is of type string
    
////---------------CONS-----------------------------------------------------
    
   const LENGTH int = 10
   const WIDTH int = 5   
   var area int

   area = LENGTH * WIDTH
   fmt.Printf("value of area : %d", area)      //  50
    
////---------------FOR-----------------------------------------------------
    
    i := 1                  // while yerine kullanabiliriz
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }                              // 1   2   3
    
    
     for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }                               // 7   8    9
    
    
    for {                        //sonsuz döngüyü break kırar
        fmt.Println("loop")
        break                        
    }                           // loop
    
    
    
     for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }                            // 1    3     5 
    
    
    
    
    
    //////-----------------İF-ELSEE----------------------------
    
    
    
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")         
    } else {
        fmt.Println(num, "has multiple digits")
    }
    
        
   ///////-----------Switch-Case----------------- 
    
    
    ii := 2
    fmt.Print("Write ", ii, " as ")
    switch ii {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }
    
  ////
    
     tt := time.Now()
    switch {
    case tt.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }
    
    
    
     
}
