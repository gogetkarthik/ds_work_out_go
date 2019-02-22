//package interview
//
//import "fmt"
//
//Given two arrays of numbers in descending order, write a function that returns a single array sorted in the same order.
//
//E.g.
//
//array1:  n
//array2: 7, 6, 5, 3 m
//
//Result array should be: 7, 6, 5, 4, 3, 2, 1
//
//asdas

//Process:
//
//An executing instance of a program is called a process.
//Some operating systems use the term ‘task‘ to refer to a program that is being executed.
//A process is always stored in the main memory also termed as the primary memory or random access memory.
//Therefore, a process is termed as an active entity. It disappears if the machine is rebooted.
//Several process may be associated with a same program.
//On a multiprocessor system, multiple processes can be executed in parallel.
//On a uni-processor system, though true parallelism is not achieved, a process scheduling algorithm is applied and the processor is scheduled to execute each process one at a time yielding an illusion of concurrency.
//Example: Executing multiple instances of the ‘Calculator’ program. Each of the instances are termed as a process.
//
//Thread:
//
//A thread is a subset of the process.
//It is termed as a ‘lightweight process’, since it is similar to a real process but executes within the context of a process and shares the same resources allotted to the process by the kernel.
//Usually, a process has only one thread of control – one set of machine instructions executing at a time.
//A process may also be made up of multiple threads of execution that execute instructions concurrently.
//Multiple threads of control can exploit the true parallelism possible on multiprocessor systems.
//On a uni-processor system, a thread scheduling algorithm is applied and the processor is scheduled to run each thread one at a time.
//All the threads running within a process share the same address space, file descriptors, stack and other process related attributes.
//Since the threads of a process share the same memory, synchronizing the access to the shared data within the process gains unprecedented importance.

//Per process items             | Per thread items
//------------------------------|-----------------
//Address space                 | Program counter
//Global variables              | Registers
//Open files                    | Stack
//Child processes               | State
//Pending alarms                |
//Signals and signal handlers   |
//Accounting information        |

package main

import "fmt"

func main() {

	arr1 := []int{4, 2, 1}

	arr2 := []int{7, 6, 5, 3}

	fmt.Println("Hello world - Go!")
	fmt.Println(MergeArrays(arr1, arr2))
}

func MergeArrays(arr1, arr2 []int) []int {

	mergeArray := make([]int, len(arr1)+len(arr2))
	var indexMerge int

	var i int
	var j int

	for {

		if arr1[i] > arr2[j] {
			mergeArray[indexMerge] = arr1[i]
			i++
		} else {
			mergeArray[indexMerge] = arr2[j]
			j++
		}

		indexMerge++

		if i == len(arr1) {
			break
		}

		if j == len(arr2) {
			break
		}
		fmt.Println(mergeArray)
	}

	fmt.Println(len(mergeArray), j, i, mergeArray[:j+i])
	if i < len(arr1) {
		mergeArray = append(mergeArray[:j+i], arr1[i:]...)
	}

	if j < len(arr2) {
		mergeArray = append(mergeArray[:j+i], arr2[j:]...)
	}
	return mergeArray
}
