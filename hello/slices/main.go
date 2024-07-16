package main

import "fmt"


func slices1() {
	s := make([]byte, 5)

	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	b := []int{1,2,3,4,5}
	c := b[2:3]

	fmt.Println("b: ", b, "c: ", c)

	// updating re-slice causes underlying array element to be updated
	c[0] = 7
	fmt.Println("b: ", b, "c: ", c)

	fmt.Println("Len(c): ", len(c), "Cap(c): ", cap(c))

	// growing the slice length to its capacity
	c = c[:cap(c)]
	fmt.Println("c: ", c)

	// copy b into c
	fmt.Println("Number of elements copied: ", copy(c, b))
	fmt.Println("c: ", c, "Len(c): ", len(c), "Cap(c): ", cap(c))

	// updating re-slice
	c[0] = 100
	fmt.Println("b: ", b, "c: ", c)

	// create a new destination slice with higher capacity
	t := make([]int, len(b), (cap(b)+1)*2)
	copy(t, b)
	fmt.Println("t: ", t, "Len(t): ", len(t), "Cap(t): ", cap(t))

	t[0] = 1000
	fmt.Println("b: ", b, "t: ", t)
	// will not affect b as t was created using make and uses a new underlying array

	/*
	Internal code of copy(dst, src)

	t := make([]byte, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
	for i := range s {
		t[i] = s[i]
	}
	s = t

	*/

	// a = make([]int, 4, 2) // Error message is "invalid argument: length and capacity swapped"

}

func AppendInt(s []int, data ...int) []int {
	fmt.Printf("%T", data) // []int
	fmt.Println("")
	currLen := len(s)
	newLen := currLen + len(data)
	if newLen > cap(s) {
		// growing the capacity of s
		t := make([]int, newLen)
		copy(t, s)
		s = t;
	}
	// now we know it has sufficient capacity, 
	// we grow it to how much we need (will be useful only when it does not go inside above if condition) 
	s = s[:newLen]
	copy(s[currLen:newLen], data)
	return s
}

func slices2() {
	// append
	a := []int{0,1}

	// lets add 3 elements
	b := make([]int, 3, 10)
	c := [3]int{100,101,102}
	copy(b, c[:])

	fmt.Println("a:", a, "len(a):", len(a), "cap(a):", cap(a))
	fmt.Println("b:", b, "len(b):", len(b), "cap(b):", cap(b))

	temp := c[:]

	c[0] = 103
	fmt.Println("c[0]:", c[0], "b[0]:", b[0], "temp[0]:", temp[0])
	// will not affect b as b was created using make which created a new underlying array
	// temp is affected cause it's slice pointer points to the same underlying array, ie. c

	aShallowCopy := a[:]
	bShallowCopy := b[:]

	a = AppendInt(a, c[0], c[1], c[2])
	b = AppendInt(b, c[:]...)

	fmt.Println("a:", a, "len(a):", len(a), "cap(a):", cap(a))
	fmt.Println("b:", b, "len(b):", len(b), "cap(b):", cap(b))

	a[0] = -1
	fmt.Println("a:", a, "aShallowCopy:", aShallowCopy, "Change not reflected as a was grown using make and now points to a different underlying array.")
	
	b[0] = -1
	fmt.Println("b:", b, "bShallowCopy:", bShallowCopy, "Change reflected as b still points to the same underlying array.")

	// ---

	// using in-built append function
	d := append(a, c[:]...)
	d = append(d, c[0], c[1], c[2])

	fmt.Println("d:", d, "len(d):", len(d), "cap(d):", cap(d))

	// ---

	e := make([]int, 3, 5)
	copy(e[:3], c[:])

	eShallowCopy := e[:]
	fmt.Println("e:", e, "eShallowCopy:", eShallowCopy)

	e = append(e, 500, 501)
	eShallowCopy = eShallowCopy[:cap(eShallowCopy)]
	fmt.Println("e:", e, "eShallowCopy:", eShallowCopy) // both are identical now

	e[0] = -1
	fmt.Println("e:", e, "eShallowCopy:", eShallowCopy, "Change reflected as both point to the same underlying array.")

	e = append(e, 502, 503) // capacity has already been reached, will use make and point to a new underlying array
	e[0] = -2
	fmt.Println("e:", e, "eShallowCopy:", eShallowCopy, "Change not reflected as both point to different underlying array.")

	// ---

	f := make([]int, 3, 5)
	copy(f[:3], c[:])

	fShallowCopy := f[:]
	fmt.Println("f:", f, "fShallowCopy:", fShallowCopy)

	f = append(f, 500, 501, 502)
	fShallowCopy = fShallowCopy[:cap(fShallowCopy)]
	fmt.Println("f:", f, "fShallowCopy:", fShallowCopy) // fShallowCopy will hold first three elements of c followed by 2 zeros
	
	// ---

}

func main() {
	slices1()
	fmt.Println("---")
	slices2()
	fmt.Println("---")
}