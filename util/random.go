package util
import(
	"math/rand"
	"time"
	"strings"
	"fmt"
)
const alphabet="abcdefghijklmnopqrstuvwxyz"

func init(){
	rand.Seed(time.Now().UnixNano())
}

// RandomeInt generates a random int between min and max
func RandomeInt(min, max int64) int64{
	return min+rand.Int63n(max-min+1) //min -max
}

// RandomeString generates a random string lenght n 
func RandomeString(n int) string{
	var sb strings.Builder
	k:= len(alphabet)

	for i := 0; i<n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates the owner name 
func RandomOwner() string{
	return RandomeString(6)
}

//  RandomMoney gernrates random money
func RandomMoney() int64{
	return RandomeInt(0,1000)
}

// RandomCurrency generates random currency
func RandomCurrency() string{
	currency :=[]string{USD,INR,EUR}
	n:= len(currency)
	return currency[rand.Intn(n)]
}

//RandomEmail generates random emails
func RandomEmail() string{
	return fmt.Sprintf("%s@email.com", RandomeString(6))
}