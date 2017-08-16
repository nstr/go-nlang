# go-nlang
Natural language processing for Go

* [Distance](#distance)
* [Phonetics](#phonetics)
* [To Do](#to-do)


## Distance

```go

import (
	"fmt"

	"github.com/njnest/go-nlang/distance"
)

func main() {
  distanceText1 := "hi"
	distanceText2 := "hello"
	distanceText3 := "hello, my darling"
	distanceTextResult := "comparison of '%s' & '%s'. Result is %f \n"

	fmt.Printf(distanceTextResult, distanceText1, distanceText2, distance.Compare(distanceText1, distanceText2))
  //comparison of 'hi' & 'hello'. Result is 0.000000
	fmt.Printf(distanceTextResult, distanceText2, distanceText3, distance.Compare(distanceText2, distanceText3))
  //comparison of 'hello' & 'hello, my darling'. Result is 0.444444
	fmt.Printf(distanceTextResult, distanceText3, distanceText3, distance.Compare(distanceText3, distanceText3))
  //comparison of 'hello, my darling' & 'hello, my darling'. Result is 1.000000
}
```

## Phonetics

```go
import (
	"fmt"

	"github.com/njnest/go-nlang/phonetics"
)

func main() {
  ph, res := phonetics.Compare("New Zeland it is part of Feladelfia", "New Zeland it is part of Pheladelphia")
  fmt.Println("Sentences sound the same:", res)
  fmt.Println("Sentences contain the same sounds:", ph.SentencesСontainTheSameWords())
}
```


## To Do
1. tokenizers
2. frequency
3. tf-idf
