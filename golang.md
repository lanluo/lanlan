<!-- TOC -->autoauto- [golang](#golang)auto    - [go types](#go-types)auto        - [string](#string)auto        - [slice and array](#slice-and-array)auto        - [map](#map)auto        - [struct](#struct)auto        - [type conversion](#type-conversion)auto        - [pass by value vs pass by pointer](#pass-by-value-vs-pass-by-pointer)auto    - [go control structure](#go-control-structure)auto    - [switch case](#switch-case)auto    - [for loop](#for-loop)autoauto<!-- /TOC -->

# golang

## go types

### string

- compare

```bash
if str == ""{
    ******
}

if len(str) <= 0{
    ****
}

str1 := "baa"
str2 := "ab"
if str1 < str2 {
    fmt.Println("less than")
} else if str1 == str2 {
    fmt.Println("equals")
} else {
    fmt.Println("greater than")
}
```

- str to byte/byte to str conversion

```bash
string ---> byte
mystr := "lanlan"
mybyte := []byte(mystr)

byte ---> string
mynewstr := string(mybyte)

get or update one byte
mybyte[2] = 'x'
b := mybyte[3]

b = 's'
fmt.Sprintf("%c", b)  ---> s
fmt.Sprintf("%b", b)  ---> 1110011
fmt.Sprintf("%d", b)  ---> 115
fmt.Sprintf("%v", b)  ---> s

substring
mystr := "hello world"
substr := mystr[1:3]  ---> "el"
substr := mystr[2:]   ---> "llo world"
substr := mystr[:4]   ---> "hell"
```

- string methods

```bash
strings.Contains(str, substr string) bool
strings.HasPrefix(str, substr string) bool
strings.HasSuffix(str, substr string) bool

strings.Index(str, substr string) int
strings.LastIndex(str, substr string) int

strings.Split(str, sep string) []string
strings.Join(a []string, sep string) string

strings.ToUpper(str string) string
strings.ToLower(str string) string

strings.TrimPrefix(str, prefix string) string
strings.TrimSuffix(str, suffix string) string
```

### slice and array

- initialize slice

```bash
myslice := []string{"lan", "luo"}

myslice := make([]string, 0)  ---->make returns the type
myslice = append(myslice, "lan")
myslice = append(myslice, "luo")

myslice := new([]string)      ----> new returns the pointer to the type
*myslice = append(*myslice, "lan")
*myslice = append(*myslice, "luo")
```

- initialize array

```bash
myarray :=[2]string{"lan", "luo"}
```

- built-in function append/copy
  - func append(slice []Type, elems ...Type) []Type
  - func copy(dest, src []Type) int
  - append/copy can only be used by slice
  - append can grow the capacity of the dst slice
  - copy cannot grow the capacity of the dst slice
  - make([]string, 3, 5)  len: 3  cap: 5

```bash
1>insert element to the end of the slice
myslice := []string{"lan", "cathy", "aria"}
myslice = append(myslice, "lucas")

2>insert element in the front of the slice
myslice := []string{"lan", "cathy", "aria"}
mynewslice := []string{"hello"}
mynewslice = append(mynewslice, myslice[:]...)

3>insert element in the middle(eg. index = 3) of the slice
index := 2
myslice := []string{"lan", "cathy", "aria"}
mynewslice := make([]string, 0)
mynewslice = append(mynewslice, myslice[:index]...)
mynewslice = append(mynewslice, "hello")
mynewslice = append(mynewslice, myslice[index:]...)
```

### map

- initialize map

```bash
mymap := make(map[string]string)

mymapPointer := new(map[string]string)
```

- add/delete key of the map

```bash
mymap :=make(make[string]bool)

mymap["lan"] = true

delete(mymap, "lan")

v, ok := mymap["lan"]
if ok{
    fmt.Printf("key: lan exists,",v)
}
```

- key of the map can be bool, string, struct(without slice or map field), pointer, array
- key of the map cannot be slice,map, struct with the slice or map field
- set implemented via map

```bash
1>create a set
myset := make(map[string]bool)

2>add an element to the set
myset["lan"] = true
myset["cathy"] = true

3> loop through the set
for k, _ :range myset{
    fmt.Println("set ele:", k)
}

4>check whether an elements is in the set
_, ok := myset["hello"]
if ok{
    fmt.Println("set ele: hello exists")
}

5>delete en element in the set
delete(myset, "lan")

```

### struct

```bash
type mystruct struct{
    name string
    age int
}
```

### type conversion

```bash
i := 115
b := (byte)(i)
s := (string)(i)
```

### pass by value vs pass by pointer

- slice,map etc is always passed by pointer(any modification will also modify the original copy)
- array, string, struct is always passed by value(if you have a large struct or array, pass by pointer of the array/struct can save the copy)


## go control structure

## switch case

'case' block in 'switch' statement break by default, it has a break implicitly,so no need to add break again

```bash
myvar := 1

switch myvar {
case 1:
    fmt.Println("this is 1")
case 2:
    fmt.Println("this is 2")
case 3, 4, 5, 6, 7, 8, 9:
    fmt.Println("this is within 3 to 9")
default:
    fmt.Println("this is out of range")
}
```

## for loop

```bash
shouldContinue := true
for shouldContinue {
    if ***{
        shouldContinue = false
    }
}

for i := 0; i < 5; i ++ {
   ****
}

for index, ele := range myslice/myarray/mystring {
   ****
}

//range over a map the order is randomized
for k, v := range mymap {
    ****
}

for index := range myslice/myarray/mystring {
    var b byte
    b = mystring[index]
}
```
