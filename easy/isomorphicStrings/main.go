package main

func isIsomorphic(s string, t string) bool {
    dict1 := map[byte]byte{}
    dict2 := map[byte]byte{}
    for i := 0; i < len(s); i++ {
        _, ok := dict1[s[i]]
        _, ok_2 := dict2[t[i]]

		if !ok {
            dict1[s[i]] = t[i]
        }

        if !ok_2 {
            dict2[t[i]] = s[i]
        }
		if s[i] != dict2[t[i]] || t[i] != dict1[s[i]]{
			return false
		}

    }

    return true
}