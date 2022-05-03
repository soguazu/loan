package utils

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"strings"
	"time"
)

// Faker struct to create functions that return's random variables
type Faker struct{}

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomUUID generates random uuid v4
func (f *Faker) RandomUUID() uuid.UUID {
	return uuid.NewV4()
}

// RandomObjectID generates random uuid v4
func (f *Faker) RandomObjectID() string {
	return primitive.NewObjectID().Hex()
}

// RandomInt generates random integers
func (f *Faker) RandomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

// RandomString generates random string
func (f *Faker) RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomName generates random names
func (f *Faker) RandomName() string {
	return f.RandomString(10)
}

// RandomWebsite generates random websites
func (f *Faker) RandomWebsite() string {
	return fmt.Sprintf("www.%v.com", f.RandomString(10))
}

// RandomType generates random type
func (f *Faker) RandomType() string {
	companyType := []string{"TECH", "NGO", "FINTECH"}
	return companyType[rand.Intn(len(companyType))]
}

// RandomFundSource generates random funding source
func (f *Faker) RandomFundSource() string {
	companyType := []string{"TECH", "NGO", "FINTECH"}
	return companyType[rand.Intn(len(companyType))]
}

// RandomNoOfEmployee generates random numbers of employee
func (f *Faker) RandomNoOfEmployee() int32 {
	return f.RandomInt(0, 10)
}

// RandomNo generates random numbers
func (f *Faker) RandomNo() int32 {
	return f.RandomInt(0, 10)
}

// RandomAddress generates random address
func (f *Faker) RandomAddress() string {
	return fmt.Sprintf("%v %v", f.RandomInt(0, 10), f.RandomName())
}
