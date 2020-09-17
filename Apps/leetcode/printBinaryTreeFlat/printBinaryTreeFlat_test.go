package printBinaryTreeFlat

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

func TestprintBinaryTreeFlat(t *Testing.t){
    for _, test := range tests{
        actual := printBinaryTreeFlat(test.input)
        if actual != test.expected{
            t.Errorf("Error. Expected: %v , Got: %v , Input: %v", test.expected, actual, test.input)
        }
    }
}
