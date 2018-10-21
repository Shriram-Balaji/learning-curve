package  sorting;

import java.util.Random;

public class QuickSort {
	public static boolean RANDOMIZE_PIVOT = true; 
    private int[] sort(int[] arrayToSort) {
        int len = arrayToSort.length;
		quickSort(arrayToSort, 0, len - 1);
        return arrayToSort;
    }

	
	private void quickSort(int[] arr, int start, int end) {
		
		// base case 
	    // exit if the partitioned segment is invalid or has only 1 element	
		if(start >= end)
			return;

		// recursive case
		int pIndex = partition(arr, start, end);
		quickSort(arr, start, pIndex - 1); // recursively sort left of pivot
		quickSort(arr, pIndex , end); // recursively sort right of pivot
	}
	// rearranges the array such that elements in the array that are lesser
	// than a chose pivot are on its left, and elements in the array that are 
	// greater than the chose pivot are on its right.
	
	// returns partitionIndex after partitioning.
	private int partition(int[] arr, int start, int end) {
        
		// pivot is the element that helps partition the array
//		int pivotIndex = findPivotIndex(start, end);
		int pivot = findPivot(arr, start, end);

		while(start <= end) {

		    /*
		     * move start pointer outwards and look for an element on the left part that is bigger than the pivot so that it can be swapped later
		     * ie. keep incrementing starting pointer until you reach a value
		       that is greater than the pivot, so that it can be swapped
		    */
		    while(arr[start] < pivot) {
		        start++;
            }

            /*
             * move end pointer inward and look for an element on the right part that is lesser than the pivot so that it can be swapped later
		     * keep decrementing ending pointer until you reach a value
		       that is lesser than the pivot, so that it can be swapped
		    */
            while(arr[end] > pivot) {
		        end--;
            }

            // now both start, and end point to the values that need to be
            // swapped because they dont satisfy the conditions in the while loop
            if(start <= end) { // this checks if start and end values are appropriate
		        swapElements(arr, start, end);
		        start++;
		        end--;
            }
        }

        // the starting pointer has been incremented till the partitionIndex
        // now elements left of the start pointer are less than the pivot
        // elements greater than pivot are on the right of start pointer
        return start;
    }

    // arr - array to swap
    private void swapElements(int[] arr, int a, int b) {
        int temp = arr[a];
        arr[a] = arr[b];
        arr[b] = temp;
    }

    private int findPivot(int arr[], int start, int end ){
        int index;
		if(RANDOMIZE_PIVOT) {
            // randomizes pivot value
            Random rand = new Random();
            int bound = end - start + 1;
            int random = rand.nextInt(bound);
            index = start + random;
        }
        else {
		    // pivot is the middle element in the array
			index = (start + end )/2;
        }
		return arr[index];
    }

    public static void main(String[] args) {
        int[] arr = {21, 12545, 23432 , 12, 35, 2000, 12313, 234};
        QuickSort qs = new QuickSort();
        int[] sorted = qs.sort(arr);
        for(int el: sorted) {
            System.out.print(el + " ");
        }
    }
}
