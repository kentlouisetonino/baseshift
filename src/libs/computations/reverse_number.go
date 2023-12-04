package computations

// This will reverse the number.
// Example is 123 to 321.
func ReverseNumber(num int) int {
   res := 0
  
   // Loop through the passed number.
   for num > 0 {
      remainder := num % 10
      res = (res * 10) + remainder
      num /= 10
   }

   return res
}
