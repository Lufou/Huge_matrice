package main

import (
	"fmt"
)

func main() {
	var tableau = []int{62, 1, 48, 37, 53, 94, 16, 84, 21, 75}
	fmt.Println(tableau)

	var tableauTrie []int = sortQuick(tableau, len(tableau), 0, len(tableau)-1)
	fmt.Println(tableauTrie)
}

func swap(d []int, a int, b int) {
  var i int
  i=d[a]
  d[a]=d[b]
  d[b]=i
}

func segmentation(d []int, N int, debut int, fin int) int {
 	var i int
 	var pivot int
 	var place int

	pivot=d[debut]; 
	place=debut;

	for i=debut+1 ; i<=fin ; i++ {
		if (d[i] <= pivot){ 
		  	place++;
		  	swap(d,i,place);
		}
	}
	swap(d, debut, place);
	return place;
}

func sortQuick(d []int, N int, debut int, fin int) []int {
	var place2 int
	if(debut<fin){
		place2 = segmentation(d, N, debut, fin);
		sortQuick(d, N, debut, place2-1);
		sortQuick(d, N, place2+1, fin);
	}
	return d
}