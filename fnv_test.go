package fnv

import "testing"

func TestVectors(t *testing.T) {
    tv_input := [...][]byte{
        []byte(""),
        []byte("a"),
        []byte("ab"),
        []byte("abc"),
        []byte("abcd"),
        []byte("abcde"),
        []byte("abcdef"),
        []byte("abcdefg"),
        []byte("abcdefgh"),
    }
    tv_output := [...]uint64{
        0xcbf29ce484222325,
        0xaf63dc4c8601ec8c,
        0x089c4407b545986a,
        0xe71fa2190541574b,
        0xfc179f83ee0724dd,
        0x6348c52d762364a8,
        0xd80bda3fbe244a0a,
        0x406e475017aa7737,
        0x25da8c1836a8d66d,
    }

    for i,x := range tv_input {
        my_output := FNV1A(x)
        expected_output := tv_output[i]
        if my_output != expected_output {
            t.Errorf("[FNV-1A] Test vector %d failed:\n\tExpected Output: %d\n\tMy Output: %d", i, expected_output, my_output)
        }
    }
}
