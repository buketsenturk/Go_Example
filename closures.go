package main
import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
// closure fonksiyon içinde fonksiyon gibi calışır



//--- recursion-------

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}



func main() {

    nextInt := intSeq()
  
    fmt.Println(nextInt())      // 1
    fmt.Println(nextInt())      // 2
    fmt.Println(nextInt())      //  3
  
    newInts := intSeq()
    fmt.Println(newInts())         //1
    
    //////-----------recursion-----------------   
    
    fmt.Println(fact(7))
    
}
  
