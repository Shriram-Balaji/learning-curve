package search;

public class BinarySearch {
    private int search(int[] arrayToSearch, int find) {
        int len = arrayToSearch.length;
        return binarySearch(arrayToSearch, find, 0, len -1 );
    }

    private int binarySearch(int[] arr, int x, int start, int end){
        if(start > end)
            return -1;
        int mid = (end + start)/2;
        if(arr[mid] < x) {
            return binarySearch(arr, x, mid + 1, end );
        }
        else if(arr[mid] > x) {
           return binarySearch(arr, x, start, mid - 1);
        }
        else
            return mid;
    }

    public static void main(String args[]) {
        BinarySearch bs = new BinarySearch();
        int[] arr = {1, 2, 3, 4};
        int foundAt = bs.search(arr, 5);
        System.out.println(foundAt);
    }
}
