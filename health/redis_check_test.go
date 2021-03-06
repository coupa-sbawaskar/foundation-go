package health

import (
	"os"
	"regexp"

	"github.com/go-redis/redis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Redis Health Checker", func() {
	Describe("Check", func() {
		if os.Getenv("TEST_REDIS") == "true" {
			It("gets Redis version with a valid client and connection", func() {
				r := redis.NewClient(&redis.Options{
					Addr:     ":6379",
					Password: "",
					DB:       0,
				})
				c := RedisCheck{Name: "testRedis", Client: r}
				d := c.Check()
				Expect(d.Name).To(Equal("testRedis"))
				Expect(d.State.Status).To(Equal(OK))
				Expect(d.Version).To(MatchRegexp(`\d+\.\d+.\d+`))
				Expect(d.Revision).To(Equal("00000000"))
			})
		}

		It("shows unknown type of Client", func() {
			c := RedisCheck{Name: "test"}
			d := c.Check()
			Expect(d.Name).To(Equal("test"))
			Expect(d.State.Status).To(Equal(CRIT))
			Expect(d.State.Details).To(Equal("Unknown type of Redis client"))
		})
	})

	Describe("getValueFromPair", func() {
		It("gets the value after the separator", func() {
			Expect(getValueFromPair(":value", ":")).To(Equal("value"))
			Expect(getValueFromPair("key:value", ":")).To(Equal("value"))
			Expect(getValueFromPair("key:value:value2", ":")).To(Equal("value:value2"))

			Expect(getValueFromPair("", ":")).To(BeEmpty())
			Expect(getValueFromPair("key:", ":")).To(BeEmpty())
			Expect(getValueFromPair("key-value", ":")).To(BeEmpty())
		})
	})

	Describe("getMatch", func() {
		It("gets the value after the separator", func() {
			re := regexp.MustCompile(`redis_version:\s*(\w|\-|\.)+`)

			Expect(getMatch("redis_version:v1.0.0", re)).To(Equal("v1.0.0"))
			Expect(getMatch("redis_version:v1.0.0 ", re)).To(Equal("v1.0.0"))
			Expect(getMatch("redis_version:v1.0.0,a", re)).To(Equal("v1.0.0"))

			Expect(getMatch("redis_version:v1.0.0alpha hi", re)).To(Equal("v1.0.0alpha"))
			Expect(getMatch("redis_version:v1.0.0-alpha\nhi", re)).To(Equal("v1.0.0-alpha"))
			Expect(getMatch("redis_version:v1.0.0_alpha", re)).To(Equal("v1.0.0_alpha"))

			Expect(getMatch("redis_version:!v1.0.0alpha", re)).To(Equal(""))
			Expect(getMatch("redis_version:,v1.0.0-alpha", re)).To(Equal(""))
		})
	})
})
