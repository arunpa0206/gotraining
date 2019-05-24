package main

import "fmt"

func main() {
   var a int
   var ptr *int
   var pptr **int

   a = 3000

   /* take the address of var */
   ptr = &a

   /* take the address of ptr using address of operator & */
   pptr = &ptr

   /* take the value using pptr */
   fmt.Printf("Value of a = %d\n", a )
   fmt.Printf("Value available at *ptr = %d\n", *ptr )
   fmt.Printf("Value available at **pptr = %d\n", **pptr)
}