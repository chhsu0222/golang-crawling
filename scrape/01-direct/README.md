# Colly

## Description

First we need a collector. Once we've got a collector, we're going to create a data structure to hold the different information we want to pull out. In this case, we want to pull out the name, username and message.
Then we use OnHTML to register a function which will be executed on every HTML element matched by the GoQuery Selector parameter, in this case we're using standard CSS selectors.
Basically, we tell the OnHTML that we're interested in everything with "tweet" class, which is a common class we can find for every tweet.
The HTMLElement in our case is a <div>. Then we use the ChildText method to get text form a sub-selected item. In this case, we're looking for a class "fullname" and one of the parents has a class of "account-group".

After we register the OnHTML function, we tell Colly to visit the site. Once we told Colly to visit, Colly actually uses go routines to speed up the visiting, so we use Wait to wait until Colly is finished scraping.
Then we've got messages full of all the tweets, so we use json.MarshalIndent to make it more readable.

Note that because Twitter does not give us all of the tweets immediately. When we scroll down to the very bottom of the page, it loads more (infinite scrolling). So we have a 2nd solution to solve that.

## References

[Colly](https://github.com/gocolly/colly)
[Reddit example](http://go-colly.org/docs/examples/reddit/)
