/*Given an array, find the average of all sub arrays of ‘K’ contiguous elements in it.*/
package main

import "fmt"

//Brute force method
func find_averages_of_subarrays(k int, arr []float64) []float64{
  var res []float64
  for i := 0; i <= len(arr) -k; i++ {
  	var sum float64 = 0.0
  	//move to the next index while grouped in 5
  	// find sum of next 'K' elements
  	for j := i; j < i + k; j++ {
  		sum += arr[j]
	}
	res = append(res, sum/float64(k))
  }
	return res
}

// Sliding window method
func find_averages_of_subarrays2(k int, arr []float64) []float64{
	var res []float64
	sum := float64(0)
	start := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if i >= k - 1 {
			res = append(res, sum/float64(k))
			sum -= arr[start]
			start += 1
		}
	}
	return res
}

func main(){
	//fmt.Println(find_averages_of_subarrays(5, []float64{1, 3, 2, 6, -1, 4, 1, 8, 2}))
	fmt.Println(find_averages_of_subarrays2(5, []float64{1, 3, 2, 6, -1, 4, 1, 8, 2}))
}