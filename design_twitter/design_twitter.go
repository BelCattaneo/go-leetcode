package main

import (
	"fmt"
	"sort"
)

	var ts int

	type User struct {
		Id int
	}

	type Tweet struct {
		Id int
		ts int
	}

	type Twitter struct {
		Tweets map[User][]Tweet
		Followers  map[User]map[User]bool
	}


	func Constructor() Twitter {
		return Twitter{
			Tweets: make(map[User][]Tweet),
			Followers: make(map[User]map[User]bool),
		}
	}


	func (this *Twitter) PostTweet(userId int, tweetId int)  {
		ts++
		this.Tweets[User{userId}] = append(this.Tweets[User{userId}], Tweet{tweetId, ts})
	}


	func (this *Twitter) GetNewsFeed(userId int) []int {
		var tweets []Tweet
		tweets = append(tweets, this.Tweets[User{userId}]...)
		for followee, _ := range this.Followers[User{userId}] {
			tweets = append(tweets, this.Tweets[followee]...)
		}

		sort.Slice(tweets, func(i, j int) bool {
			return tweets[i].ts > tweets[j].ts
		})

		var feedIds []int
		for i, tweet := range tweets {
			feedIds = append(feedIds, tweet.Id)
			if i == 9 {
				return feedIds
			}
		}
		return feedIds
	}


	func (this *Twitter) Follow(followerId int, followeeId int)  {
		follower := User{followerId}
		followee := User{followeeId}
		if this.Followers[follower] == nil {
			this.Followers[follower] = make(map[User]bool)
		}
		this.Followers[follower][followee] = true
	}


	func (this *Twitter) Unfollow(followerId int, followeeId int)  {
		follower := User{followerId}
		followee := User{followeeId}
		delete(this.Followers[follower], followee)
	}


func main() {
	twitter := Constructor()
	twitter.PostTweet(1, 5) // User 1 posts a new tweet (id = 5).
	twitter.PostTweet(1, 3) // User 1 posts a new tweet (id = 5).
	twitter.PostTweet(1, 101) // User 1 posts a new tweet (id = 5).
	twitter.PostTweet(1, 13) // User 1 posts a new tweet (id = 5).
	twitter.PostTweet(1, 10) // User 1 posts a new tweet (id = 5).
	twitter.PostTweet(1, 2) // User 1 posts a new tweet (id = 5).
	twitter.PostTweet(1, 94) // User 1 posts a new tweet (id = 5).
	twitter.PostTweet(1, 505) // User 1 posts a new tweet (id = 5).
	twitter.PostTweet(1, 333) // User 1 posts a new tweet (id = 5).
	twitter.PostTweet(1, 22) // User 1 posts a new tweet (id = 5).
	twitter.PostTweet(1, 11) // User 1 posts a new tweet (id = 5).
	//fmt.Println(twitter.Tweets)
	twitter.GetNewsFeed(1)  // User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
	//fmt.Println(feed1)
	twitter.Follow(1, 2)   // User 1 follows user 2.
	twitter.PostTweet(2, 6) // User 2 posts a new tweet (id = 6).
	twitter.GetNewsFeed(1)  // User 1's news feed should return a list with 2 tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
	twitter.Unfollow(1, 2)  // User 1 unfollows user 2.
	feed2 := twitter.GetNewsFeed(1)  // User 1's news feed should return a list with 1 tweet id -> [5], since user 1 is no longer following user 2.
	fmt.Println(feed2)
	/*
	fmt.Println(param2)
	*/
}
