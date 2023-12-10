package computations

func NumberArray(num int64) []int64 {
  var arrayOfNum []int64
  
  for num != 0 {
    arrayOfNum = append(arrayOfNum, num % 10)
    num /= 10
  } 

  return arrayOfNum
}
