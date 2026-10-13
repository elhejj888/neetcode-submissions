func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1Pos := m-1
	nums2Pos := n-1
	tale := len(nums1) - 1

	for nums2Pos >= 0 {
		if nums1Pos >= 0 && nums1[nums1Pos] > nums2[nums2Pos] {
			nums1[tale] = nums1[nums1Pos]
			nums1Pos--
		} else {
			nums1[tale] = nums2[nums2Pos]
			nums2Pos--
		}
		tale--
	}
}
