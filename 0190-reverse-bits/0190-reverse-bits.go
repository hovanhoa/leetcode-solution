func reverseBits(num uint32) uint32 {
    var answer uint32
    for i := 0; i < 32 ; i++{
        answer=answer | (num&1)<<(31-i)
        num=num>>1
    }
    return answer 
}