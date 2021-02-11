/*
You are given a string made of lowercase English letters a, b, and c. Your task is to minimize the length of the string by applying the following operation on the string:

Divide the string into two non-empty parts, left and right part.
Without reversing any of the parts, swap them with each other by appending the left part to the end of the right part.
While appending, if the suffix string of the right part and the prefix string of the left part contain the same character, then you can remove those characters from the suffix and prefix of the right and left part respectively.
Repeat the third step until you do not find such prefix and suffix strings
Determine the minimum length of the string after applying the above operations exactly once on a string.

For example, you are given the following string: aabcccabba

Step 1: left part: aabcc and right part: cabba

Appending both strings:

In the string cabbaaabcc, you can remove aaa from the string and the new string will be cabbbcc.

In the string cabbbcc, you can remove bbb from the string and the new string will be cacc.

You cannot reduce the string cacc.

Input format

First line:

Output format

Print a single integer denoting the minimum length of the string after applying the operation.
*/

package main

import (
	"fmt"
)

func main() {

	var str string
	//fmt.Scanf("%s", &str)
	str = "aabcccabba"
	l := len(str)
	ans := minimize(str,l)
    fmt.Println(ans)
}

func minimize(s string,l int) int  {
	leftPart := s[0:(l/2)] //aabcc
	rightPart := s[(l/2):] //cabba
	if leftPart[0] != rightPart[len(rightPart)-1]{
      return l
	}

	//combine - cabba abbcc
	i,j,count := len(rightPart)-1,0,0
	matched := true
	for (i >= 0 || j <= len(rightPart)) && matched == true{
		if rightPart[i] == leftPart[j]{
			count = count+2
			for rightPart[i] == rightPart[i-1]{
				count++
				i--
			}
			for leftPart[j] == leftPart[j+1]{
				count++
				j++
			}
			i--
			j++
		}else{
			i--
			j++
			matched = false
		}
	}
	return l-count
}
