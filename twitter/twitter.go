package twitter

import (
	"sort"
)

// Twitter System
type Twitter struct {
	// Time
	timeMoment int

	// tweetId --> [0] authorId, [1] timeMomentId
	tweetMoment map[int]int

	// followeeId --> followersId
	follower map[int]map[int]bool

	// followerId --> followeeIds
	followee map[int]map[int]bool

	// authorId --> tweetsIds
	userTweets map[int][]int
}

// Twitter methods
////////////////////////////////////////////////////////////////////////////////////////////////////

func Constructor() Twitter {
	return Twitter{
		timeMoment:  0,
		tweetMoment: make(map[int]int),
		follower:    make(map[int]map[int]bool),
		followee:    make(map[int]map[int]bool),
		userTweets:  make(map[int][]int)}
}

func (t *Twitter) checkIn(userId int) {
	if _, registered := t.follower[userId]; !registered {
		t.follower[userId] = map[int]bool{}
		t.followee[userId] = map[int]bool{}
		t.userTweets[userId] = make([]int, 0)
	}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.checkIn(userId)
	t.tweetMoment[tweetId] = t.timeMoment
	t.timeMoment++
	t.userTweets[userId] = append(t.userTweets[userId], tweetId)
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	t.checkIn(followerId)
	t.checkIn(followeeId)
	t.follower[followeeId][followerId] = true
	t.followee[followerId][followeeId] = true
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	delete(t.follower[followeeId], followerId)
	delete(t.followee[followerId], followeeId)
}

func arrayTail(array []int, n int) []int {
	k := len(array)
	if k < n {
		return array
	}
	return array[k-n:]
}

func arrayHead(array []int, n int) []int {
	k := len(array)
	if k < n {
		return array
	}
	return array[:n]
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	t.checkIn(userId)
	feed := arrayTail(t.userTweets[userId], 10)
	for followeeId := range t.followee[userId] {
		feed = append(feed, arrayTail(t.userTweets[followeeId], 10)...)
	}
	sort.Slice(feed, func(i, j int) bool {
		return t.tweetMoment[feed[i]] > t.tweetMoment[feed[j]]
	})
	return arrayHead(feed, 10)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
