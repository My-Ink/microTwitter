package tests

import (
	"Twitter/twitter"
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Log("Assertion failed:\n[LOG] Expected: ", expected, "\n[LOG] Got: ", actual)
		t.FailNow()
	}
}

func TestPostTweet(t *testing.T) {
	t.Run("Post Mechanics", func(t *testing.T) {
		obj := twitter.Constructor()
		obj.PostTweet(0, 3)
		obj.PostTweet(0, 2)
		obj.PostTweet(100, 18)
		AssertEqual(t, obj.GetNewsFeed(0), []int{2, 3})
		AssertEqual(t, obj.GetNewsFeed(100), []int{18})
	})
}

func TestFollow(t *testing.T) {
	t.Run("Post Before Follow", func(t *testing.T) {
		obj := twitter.Constructor()
		obj.PostTweet(1, 5)
		obj.Follow(2, 1)
		AssertEqual(t, obj.GetNewsFeed(2), []int{5})
	})

	t.Run("Post After Follow", func(t *testing.T) {
		obj := twitter.Constructor()
		obj.Follow(2, 1)
		obj.PostTweet(1, 5)
		AssertEqual(t, obj.GetNewsFeed(2), []int{5})
	})
}

func TestUnfollow(t *testing.T) {
	t.Run("Follow-Unfollow", func(t *testing.T) {
		obj := twitter.Constructor()
		obj.PostTweet(1, 1)
		obj.Follow(2, 1)
		AssertEqual(t, obj.GetNewsFeed(2), []int{1})
		obj.Unfollow(2, 1)
		AssertEqual(t, obj.GetNewsFeed(2), []int{})
	})
}

func TestStress(t *testing.T) {
	t.Run("Sequential Stress Test", func(t *testing.T) {
		obj := twitter.Constructor()
		for i := 0; i < 10; i += 2 {
			for j := 0; j < 5; j++ {
				obj.PostTweet(i, 5*i+j)
				obj.PostTweet(i+1, 5*i+5+j)
			}
			obj.Follow(i, i+1)
			obj.Follow(i+1, i+2)
		}
		for i := 1; i < 10; i += 2 {
			obj.Unfollow(i, i+1)
		}

		for i := 0; i < 10; i++ {
			expected := make([]int, 0)
			if i%2 == 1 {
				for j := 0; j < 5; j++ {
					expected = append(expected, 5*i+4-j)
				}
			} else {
				for j := 0; j < 5; j++ {
					expected = append(expected, 5*i+9-j)
					expected = append(expected, 5*i+4-j)
				}
			}
			AssertEqual(t, obj.GetNewsFeed(i), expected)
		}
	})
}
