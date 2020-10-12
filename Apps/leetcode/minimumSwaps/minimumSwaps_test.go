package minimumSwaps

import "testing"

var tests = []struct{
    //var input
    //var expected
}{
  /*
    {},
    {},
    {},
    {}, 
    {},
    {},
  */
}

func TestminimumSwaps(t *Testing.t){
    for _, test := range tests{
        actual := minimumSwaps(test.input)
        if actual != test.expected{
            t.Errorf("Error. Expected: %v , Got: %v , Input: %v", test.expected, actual, test.input)
        }
    }
}
