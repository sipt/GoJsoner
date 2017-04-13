package GoJsoner

import (
	"bytes"
	"errors"
)

//Map alias map[string]string
type Map map[string]string
type rMap map[rune]interface{}

//Maches discard from "start" to "end"
var Maches = []Map{
	Map{"start": "//", "end": "\n"},
	Map{"start": "/*", "end": "*/"},
}

//Discard discarding comments
//@params:content
//@resturns: string simple json
func Discard(content string) (string, error) {
	var (
		buffer    bytes.Buffer
		flag      int
		v         rune
		protected bool
	)
	runes := []rune(content)
	flag = -1
	for i := 0; i < len(runes); {
		v = runes[i]
		if flag == -1 {
			//match start
			for f, v := range Maches {
				l := match(&runes, i, v["start"])
				if l != 0 {
					flag = f
					i += l
					break
				}
			}
			if flag == -1 {
				if protected {
					buffer.WriteRune(v)
					if v == '"' {
						protected = true
					}
				} else {
					r := filter(v)
					if r != 0 {
						buffer.WriteRune(v)
					}
				}
			} else {
				continue
			}
		} else {
			//match end
			l := match(&runes, i, Maches[flag]["end"])
			if l != 0 {
				flag = -1
				i += l
				continue
			}
		}
		i++
	}
	return buffer.String(), nil
}

func filter(v rune) rune {
	switch v {
	case ' ':
	case '\n':
	case '\t':
	default:
		return v
	}
	return 0
}

func match(runes *[]rune, i int, dst string) int {
	dstLen := len([]rune(dst))
	// fmt.Println("dstLen:", dstLen, ", index:", i, ",runesLen:", len(*runes))
	// fmt.Println(string((*runes)[i : i+dstLen]))
	if len(*runes)-i >= dstLen && string((*runes)[i:i+dstLen]) == dst {
		return dstLen
	}
	return 0
}

//Stack rune stack
type Stack []rune

//Push stack push
func (s Stack) Push(r rune) {
	s = append(s, r)
}

//Pop stack pop
func (s Stack) Pop() (rune, error) {
	if len(s) == 0 {
		return 0, errors.New("stack is empty")
	}
	v := s[len(s)-1]
	s = s[:len(s)-1]
	return v, nil
}
