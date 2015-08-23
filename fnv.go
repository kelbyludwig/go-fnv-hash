package fnv

func FNV1A(input []byte) uint64 {
    var hash uint64 = 0xcbf29ce484222325
    var fnv_prime uint64 = 0x100000001b3
    for _,b := range input {
        hash ^= uint64(b)
        hash *= fnv_prime
    }
    return hash
}
