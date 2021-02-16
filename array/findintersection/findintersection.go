package findintersection
import (
  "strings"
)

func splitStrArr(str string) []string {
  splitted := strings.Split(str, ",")

  var r []string
  for _,v := range splitted {
    r = append(r, strings.Trim(v, " "))
  }

  return r
}

func buildMap(strArr []string) map[string]int {
  m := make(map[string]int)

  for _, v := range strArr {
    m[v]++
  }

  return m
} 

// FindIntersection returns a string with common elements
func FindIntersection(strArr []string) string {
  splitted := splitStrArr(strArr[0])
  m1 := buildMap(splitted)
  m2 := buildMap(splitStrArr(strArr[1]))
 
  
  var r []string
  for _, v := range splitted {
    if (m1[v] > 0 && m2[v] > 0) {
      if (m1[v] <= m2[v]) {
        for i := 0; i < m1[v]; i++ {
          r = append(r, v)
        }
      } else {
        for i := 0; i < m2[v]; i++ {
          r = append(r, v)
        }
      }
    }
  }

  if (len(r) == 0) {
    return "false"
  }

  return strings.Join(r, ",");

}