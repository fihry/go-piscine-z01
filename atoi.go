package piscine
func Atoi(s string) int {
  res:=0
  sign:=1
  for i:=0; i<=len(s)-1;i++{
    if s[0]=='-'{
      sign *= -1 
      i++
    }else if s[i]=='+'{
      i++
    }
    if s[i]>='1' && s[i]<='9'{
      res = res*10 + int(s[i]-'0')
    }else{
      return 0
    }
  }
  return res*sign
}