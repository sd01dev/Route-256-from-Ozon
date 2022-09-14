package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	value int
	left  int
	right int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var num int
	fmt.Fscan(in, &num)

	for k := 1; k <= num; k++ {
		var tmp int
		var elements int
		fmt.Fscan(in, &elements)

		var value, firstValue, secondValue int
		var nodeArray []node
		fmt.Fscan(in, &value, &firstValue, &secondValue)
		nodeArray = append(nodeArray, node{
			value: value,
			left:  firstValue,
			right: secondValue,
		})

		var unsorted []node

		for i := 0; i < elements-1; i++ {
			fmt.Fscan(in, &value, &firstValue, &secondValue)

			for _, currentNode := range nodeArray {

				if currentNode.left == value {

					if currentNode.value == firstValue {
						tmp = secondValue
					} else {
						tmp = firstValue
					}

					nodeArray = append(nodeArray, node{
						value: value,
						left:  tmp,
						right: currentNode.value,
					})
					break
				}

				if currentNode.right == value {

					if currentNode.value == firstValue {
						tmp = secondValue
					} else {
						tmp = firstValue
					}

					nodeArray = append(nodeArray, node{
						value: value,
						left:  currentNode.value,
						right: tmp,
					})
					break
				}

				if currentNode.left == firstValue {
					nodeArray = append(nodeArray, node{
						value: value,
						left:  secondValue,
						right: firstValue,
					})
					break
				}

				if currentNode.right == firstValue {
					nodeArray = append(nodeArray, node{
						value: value,
						left:  firstValue,
						right: secondValue,
					})
					break
				}

				if currentNode.right == secondValue {
					nodeArray = append(nodeArray, node{
						value: value,
						left:  secondValue,
						right: firstValue,
					})
					break
				}

				if currentNode.left == secondValue {
					nodeArray = append(nodeArray, node{
						value: value,
						left:  firstValue,
						right: secondValue,
					})
					break
				} else {

					if currentNode == nodeArray[len(nodeArray)-1] {
						unsorted = append(unsorted, node{
							value: value,
							left:  firstValue,
							right: secondValue,
						})
					}
				}
			}
		}

		for len(nodeArray) != elements {
			for i := 0; i < len(unsorted); i++ {

				for _, currentNode := range nodeArray {

					if currentNode.left == unsorted[i].value {

						if currentNode.value == unsorted[i].left {
							tmp = unsorted[i].right
						} else {
							tmp = unsorted[i].left
						}

						nodeArray = append(nodeArray, node{
							value: unsorted[i].value,
							left:  tmp,
							right: currentNode.value,
						})

						break
					}

					if currentNode.right == unsorted[i].value {

						if currentNode.value == unsorted[i].left {
							tmp = unsorted[i].right
						} else {
							tmp = unsorted[i].left
						}

						nodeArray = append(nodeArray, node{
							value: unsorted[i].value,
							left:  currentNode.value,
							right: tmp,
						})

						break
					}

					if currentNode.left == unsorted[i].left {
						nodeArray = append(nodeArray, node{
							value: unsorted[i].value,
							left:  unsorted[i].right,
							right: currentNode.left,
						})
						break
					}

					if currentNode.right == unsorted[i].left {
						nodeArray = append(nodeArray, node{
							value: unsorted[i].value,
							left:  currentNode.right,
							right: unsorted[i].right,
						})
						break
					}

					if currentNode.right == unsorted[i].right {
						nodeArray = append(nodeArray, node{
							value: unsorted[i].value,
							left:  currentNode.right,
							right: unsorted[i].left,
						})
						break
					}

					if currentNode.left == unsorted[i].right {
						nodeArray = append(nodeArray, node{
							value: unsorted[i].value,
							left:  unsorted[i].left,
							right: currentNode.left,
						})
						break
					}
				}
			}
		}

		var startPoint, startRightPoint, currentPoint, rightPoint int
		next := true
		starterNode := nodeArray[0]
		startPoint = starterNode.value
		startRightPoint = starterNode.right

		rightPoint = startRightPoint
		for n := 0; n < elements/2; n++ {

			pointToPrint := startPoint
			rightPoint = startRightPoint

			for j := 0; j < elements/2; j++ {
				for _, currentNode := range nodeArray {

					if rightPoint == currentNode.value {
						rightPoint = currentNode.right
						currentPoint = currentNode.value

						if next {
							startPoint = currentNode.value
							startRightPoint = currentNode.right
							next = false
						}

						break
					}
				}
			}
			fmt.Fprintf(out, "%d %d \n", pointToPrint, currentPoint)
			next = true
		}
	}
}
