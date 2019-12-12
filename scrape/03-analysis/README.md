# Analysis - word count

## Target

Count how many times each word has appeared in each of the tweets.

## Description

1. Create a dumpTweets function which is responsible for printing tweets in JSON format.
2. Create a data structure, wordInfo, that can be sorted for the count of a word. So we can have the word counts in descending or ascending order.
3. Create wordCount function to actually count the words in a slice of tweets. While counting words, a map is useful.
4. Loop each tweet and split it into words. Then loop each word and count how many times it appeared. Note that we use strings.ToLower so it's case insensitive. One more thing need to be noted is that the default value (zero value) of our map (int) is 0 so we can add 1 to it directly.
5. Convert the map into a slice of wordInfo, which contains the word and the count.
6. Sort the slice with sort.Slice in descending order (count).

## References

[sort.Slice](https://golang.org/pkg/sort/#Slice)
[strings.Split](https://golang.org/pkg/strings/#Split)
[strings.ToLower](https://golang.org/pkg/strings/#ToLower)
[zero value](https://golang.org/ref/spec#The_zero_value)
