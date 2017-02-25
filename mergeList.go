
/*Write a program that takes two sorted lists of integers and finds the K'th smallest element in the union of the two lists.
The input comprises of an integer M, followed by a list of M integers on the next line, then an integer N, followed by a list of N integers on the next line. The last line of the input is the integer K.
Your program should output the K'th smallest element in the two lists.
*/
package main

import "os"
import "strings"
import "fmt"
import "bufio"
import "strconv"

func main() {
  
  reader := bufio.NewReader(os.Stdin)
  
  l1,_  := reader.ReadString('\n')
  l1 = strings.Replace(l1,"\n","",-1)
  l1 = strings.Trim(l1,"\n")
  a1,_ := reader.ReadString('\n')
  arr1 := strings.Split(a1," ")
 
   
  l2,_ := reader.ReadString('\n')
  l2 = strings.Replace(l2,"\n","",-1)
  a2,_ := reader.ReadString('\n')
  arr2 := strings.Split(a2," ")
  
  kTerm,_ := reader.ReadString('\n')
  kTerm = strings.Replace(kTerm,"\n","",-1)
  

  
  firstLength := len(arr1)
  secondLength := len(arr2)
 
  kT,_ := strconv.Atoi(kTerm)
  
  //uLength := firstLength +secondLength 
  
  var firstArr = convertStrToInt(firstLength,arr1)
  var secondArr = convertStrToInt(secondLength,arr2)
  var finalMerge = mergeArr(firstArr,secondArr)
  
  
  var sortedArr = mergeSort(finalMerge)
  fmt.Println(sortedArr)  
  fmt.Println(sortedArr[kT-1])
}

func convertStrToInt(length int, arr []string)[]int{
   var intArr = make([]int,length)
   var i=0
    for i < length {
    arr[i] = strings.Replace(arr[i],"\n","",-1)
    element,err := strconv.Atoi(arr[i])
    
     if( err != nil){
      panic(err)
     }
    intArr[i] = element
    i++
  }
  return intArr
}

func mergeArr(first []int,second []int)[]int{
  
  var finLength = len(first) + len(second)
 
  var mergeArr = make([]int,finLength)
  var i = 0
  var j = 0
  for i < len(first) {
    mergeArr[i] =first[i]
    i++
    j=i
  }
  i=0
  for i < len(second) {
    mergeArr[j] =second[i]
    i++
    j++
  }
 
  return mergeArr
}

func  mergeSort(finArr []int)[]int{
  if(len(finArr) <= 1){
    return finArr
  }

    var mid = len(finArr) /2 
    var left = mergeSort(finArr[:mid])
    var right = mergeSort(finArr[mid:])
    return sorter(left,right)
}

func sorter(left , right []int)[]int{
  
  sortedArray := make([]int,0,len(left)+len(right))
  
  for (len(left) > 0 || len(right) > 0){
      if len(left) == 0 {
      return append(sortedArray,right...)
   }
    if len(right) == 0{
          return append(sortedArray,left...)
    }
    if left[0] <= right[0] {
      
      sortedArray = append(sortedArray,left[0])
      left = left[1:]
     }else{
      sortedArray = append(sortedArray,right[0])
      right =right[1:]
    } 
  }
 return sortedArray
} 
