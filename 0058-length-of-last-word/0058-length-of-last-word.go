func lengthOfLastWord(s string) int {
    s = strings.TrimSpace(s)
    sarr := strings.Split(s," ")
    return len(sarr[len(sarr)-1])
}