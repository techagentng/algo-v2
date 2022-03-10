/*Given an array, find the average of all sub arrays of ‘K’ contiguous elements in it.*/
package main

import "fmt"

//Brute force method
func find_averages_of_subarrays(k int, arr []float64) []float64{
  var res []float64
  for i := 0; i <= len(arr) -k; i++ {
  	var sum float64 = 0.0
  	//move to the next index while grouped in 5
  	for j := i; j < i + k; j++ {
  		sum += arr[j]
	}
	res = append(res, sum/float64(k))
  }
	return res
}
func main(){
	fmt.Println(find_averages_of_subarrays(5, []float64{1, 3, 2, 6, -1, 4, 1, 8, 2}))
}