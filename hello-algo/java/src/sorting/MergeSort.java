package sorting; 

public class MergeSort {
	
	private int[] sort(int[] arrayToSort) {
		int len = arrayToSort.length;
		// create temporary array to store intermediate sorted values
		int[] tmp = new int[len];
		doSort(arrayToSort, tmp, 0, len-1);
		return arrayToSort;
	}
	
	private static void doSort(int[] arr, int[] tmp, int leftStart, int rightEnd) {
		// arr - array to be sorted
		// leftStart - left most index
		// rightEnf - rightmost index
		// tmp - temporary array

        // base case
		if(leftStart >= rightEnd) {
            return;
        }

        // recursive case
        // right + left / 2 would overflow for integers whose sum exceeds/overflows the max allowed integer, causing a negative value leading to ArrayOutOfBoundsException
        int mid = leftStart + (rightEnd - leftStart)/2;
        doSort(arr, tmp, leftStart, mid); // recursively sort left half
        doSort(arr, tmp, mid + 1, rightEnd); // recursively sort right half
        mergeHalves(arr, tmp, leftStart, rightEnd); // merge the two sorted halves
	    }
	
	// merges the two sorted halves based on comparing each of them
	private static void mergeHalves(int[] arr, int[] tmp, int leftStart,  int rightEnd) {
	
		int leftEnd = (leftStart + rightEnd)/2;
		int rightStart = leftEnd + 1;
	
		// initialize pointer index variables for comparison of elements in both halves
		int left = leftStart;
		int right = rightStart;
		// current is the index of the while loop
		int current = leftStart;

		// while iterating within bounds of the array
		while(left <= leftEnd && right <= rightEnd) {
			// if element in left half is less than or equal element in right half
			if(arr[left] <= arr[right])  {
				// copy the left element into the temporary array
				tmp[current] = arr[left];
				// increment left index pointer
				left++;
			}
			else {
				// copy the right element into the temporary array
				tmp[current] = arr[right];
				// increment right pointer and move right
				right++;
			}
			// increment current pointer
			current++;
		}
		
		// System.arraycopy copies elements of a source array to a destination array where we can provide specific boundaries for copying
		// Syntax: System.arraycopy(sourceArray, sourceIndexStart, destinationArray, destinationIndexStart, numberOfElementsToCopy)
	    int remLeft = leftEnd - left + 1;
		int remRight = rightEnd - right + 1;
		int totalSize = rightEnd - leftStart + 1;

		// copy over remaining elements, once the while exits ( after reaching out of bounds )
		System.arraycopy(arr, left, tmp, current, remLeft);
		// if the left half has remaining elements to copy ,the value of left pointer ( incremented in each while iteration) will indicate the
        // index to copy from

		System.arraycopy(arr, right, tmp, current, remRight);
		// if the right half has remaining elements to copy ,the value of right pointer ( incremented in each while iteration) will indicate the
        // index to copy from


		// copy tmp array back into original array
		System.arraycopy(tmp, leftStart, arr, leftStart, totalSize);
	}
	
	public static void main(String[] args) {
		int[] arr = {3, 8, 2, 9, 10, 100, 12000, 45};
		MergeSort mergeSort = new MergeSort();
		int[] sorted = mergeSort.sort(arr);
        for (int element: sorted) {
            System.out.print(element + " ");
        }
	}
}
