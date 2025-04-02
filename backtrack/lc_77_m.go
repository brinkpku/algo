package backtrack

func combine(n int, k int) [][]int {
    var backtrack func(s int)
    ans := [][]int{}
    track := []int{}
    backtrack = func(s int){
        if len(track) == k{
            tmp := make([]int, k)
            copy(tmp, track)
            ans = append(ans, tmp)
            return
        }
        for i := s; i<=n;i++{
            track = append(track, i)
            backtrack(i+1)
            track = track[:len(track)-1]
        }
    }
    backtrack(1)
    return ans
}