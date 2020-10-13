package rotateLeft

func rotLeft(a []int32, d int32) []int32 {
  l := len(a)
  actualLeftMoves := d % int32(l)
  m := make(map[int]int32)
  for i,x := range a {
    in := i - int(actualLeftMoves)
    if in < 0 {
      in = in + l
    }
    m[in] = x
  }

  for i := 0; i < l; i++{
    a[i] = m[i]
  }
  return a
}
