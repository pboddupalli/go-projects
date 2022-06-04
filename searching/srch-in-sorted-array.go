package searching

func srchInSortedArray(nums []int, target int) int {
  i, j := 0, len(nums)
  for i < j {
    mid := i + (j - i) / 2
    if target == nums[mid] {
      return mid
    }
    if nums[i] <= nums[j-1] {
      // case of binary searchInternal
      if target < nums[mid] {
        j = mid
      } else {
        i = mid + 1
      }
    } else if nums[mid] <= nums[j-1] {
      // right half is monotonically increasing
      if target > nums[mid] && target <= nums[j-1] {
        i = mid + 1
      } else {
        j = mid
      }
    } else {
      // left half is monotonically increasing
      if target >= nums[i] && target < nums[mid] {
        j = mid
      } else {
        i = mid + 1
      }
    }
  }
  return -1
}

func search(nums []int, target int) bool {
  return searchInternal(0, len(nums), nums, target)
}

func searchInternal(i, j int, nums []int, target int) bool {
  if i >= j {
    return false
  }
  if j - i == 1 {
    // single element limiting case
    return nums[i] == target
  }
  mid := i + (j - i) / 2
  if target == nums[mid] {
    return true
  }
  if nums[mid] < nums[j-1] && target > nums[mid] && target <= nums[j-1] {
    return searchInternal(mid+1, j, nums, target)
  } else if nums[i] < nums[mid] && target >= nums[i] && target < nums[mid] {
    return searchInternal(i, mid, nums, target)
  } else {
    return searchInternal(i, mid, nums, target) || searchInternal(mid+1, j, nums, target)
  }
}

func searchVer2(nums []int, target int) bool {
  i, j := 0, len(nums)
  for i < j {
    mid := i + (j - i) / 2
    if target == nums[mid] {
      return true
    }
    if mid+1 < j && nums[mid+1] < nums[j-1] && target > nums[mid] && target <= nums[j-1] {
      // right half is increasing sequence
      i = mid + 1
    } else if i < mid && nums[i] < nums[mid] && target >= nums[i] && target < nums[mid] {
      j = mid
    } else {
      if nums[i] == target {
        return true
      }
      i = i + 1
    }
  }
  return false
}

func searchVer3(nums []int, target int) bool {
  left, right := 0, len(nums)-1

  for left <= right {
    mid := left + (right-left)/2
    if target == nums[mid] {
      return true
    }
    if (nums[left] == nums[mid]) && (nums[right] == nums[mid]){
      left++;
      right--;
    } else if nums[left] <= nums[mid] {
      if target >= nums[left] && target <=nums[mid] {
        right = mid-1
      } else {
        left = mid+1
      }
    } else {
      if target >= nums[mid] && target <=nums[right] {
        left = mid+1
      } else {
        right = mid-1
      }
    }
  }
  return false
}