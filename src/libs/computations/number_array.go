package computations

func NumberArray(num int) []int {
  var arrayOfNum []int
  
  for num != 0 {
    arrayOfNum = append(arrayOfNum, num % 10)
    num /= 10
  } 

  return arrayOfNum
}
