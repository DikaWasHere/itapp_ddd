package pkg
  
// PadLeft pads a string with leading zeros to reach a minimum length of 3
func PadLeft(s string) string {
    for len(s) < 3 {
        s = "0" + s
    }
    return s
}