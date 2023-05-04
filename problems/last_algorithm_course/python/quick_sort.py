def partition(array, low, high):
    pivot = array[high]
    
    index = low - 1

    for i in range(low, high):
        if array[i] <= pivot:
            index += 1
            temp = array[i]
            array[i] = array[index]
            array[index] = temp

    index += 1
    array[high] = array[index]
    array[index] = pivot

    return index
        

def quick_sort(array, low, high):
    if low >= high:
        return

    pivot_index = partition(array, low, high)
    quick_sort(array, low, pivot_index - 1)
    quick_sort(array, pivot_index + 1, high)
    return array

array = [0, 100, 4, 200, 2, 1]

print(quick_sort(array, 0, len(array)-1))