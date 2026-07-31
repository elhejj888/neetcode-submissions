func shipWithinDays(weights []int, days int) int {
    left := 0
    right := 0

    // Define the search space.
    for _, weight := range weights {
        if weight > left {
            left = weight
        }
        right += weight
    }

    for left < right {
        mid := left + (right-left)/2

        if canShip(weights, days, mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }

    return left
}

func canShip(weights []int, days int, capacity int) bool {
    currentLoad := 0
    usedDays := 1

    for _, weight := range weights {
        if currentLoad+weight > capacity {
            usedDays++
            currentLoad = 0
        }

        currentLoad += weight
    }

    return usedDays <= days
}