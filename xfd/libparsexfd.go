package main

import (
	"C"
	"strings"
	"strconv"
	"fmt"
	"encoding/json"
)

type Xfd struct {
	Name string		`json:"name"`
	Offset int		`json:"offset"`
	Length int      `json:"length"`
	Type int		`json:"type"`
	Digits int		`json:"digits"`
	Scale int		`json:"scale"`
	UserType int	`json:"usertype"`
	Condition int	`json:"condition"`
	Level int		`json:"level"`
	Format string	`json:"format"`
}

var xfdp map[int]Xfd

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		a := []byte(s)
		switch a[len(a)-1] {
		case 125:
			a[len(a)-1] = 48
		case 74:
			a[len(a)-1] = 49
		case 75:
			a[len(a)-1] = 50
		case 76:
			a[len(a)-1] = 51
		case 77:
			a[len(a)-1] = 52
		case 78:
			a[len(a)-1] = 53
		case 79:
			a[len(a)-1] = 54
		case 80:
			a[len(a)-1] = 55
		case 81:
			a[len(a)-1] = 56
		case 82:
			a[len(a)-1] = 57
		}
		s = string(a)
		i, err := strconv.Atoi(s)
        	if err != nil {
			return 0
		}
		return i * -1
	}
	return i
}

/*XFDParse is a func to parse a XFD*/
//export XFDParse
func XFDParse(str *C.char, ret *int) {
	strl := C.GoString(str)
	strs := strings.Split(strl, ",")
	i := len(xfdp)
	if i == 0 {
		xfdp = make(map[int]Xfd)
	}
	xfdp[i] = Xfd{
		    Name: strings.TrimSpace(strs[0]), 
            Offset: atoi(strs[1]),
            Length: atoi(strs[2]),
            Type: atoi(strs[3]),
            Digits: atoi(strs[4]),
            Scale: atoi(strs[5]),
            UserType: atoi(strs[6]),
            Condition: atoi(strs[7]),
            Level: atoi(strs[8]),
            Format: strings.TrimSpace(strs[9])}
	*ret = i
	return
}

/*XFDtoJson is a func to convert XFD to JSON*/
//export XFDtoJson
func XFDtoJson(xfdjson **C.char, length *int) {
	j, _ := json.MarshalIndent(xfdp, "", "    ")
	var reta string
	reta = fmt.Sprint(string(j))
	fmt.Println(reta)
	fmt.Println(len(reta))
	*xfdjson = C.CString(reta)
	*length = len(reta)
}

func main() {}