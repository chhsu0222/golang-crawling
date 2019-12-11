# goquery

## Check robots.txt

Because go-Colly only works on full HTML pages, so we could not use it. If we want to be respectful to the site, we have to obey the robots.txt ourselves (go-Colly takes care robots.txt for us).

We can copy the content of the robots.txt and paste it on robots.txt validator. Then we paste the URL (without the query parameters) we want to crawl and select the User agent, in our case is All.

## Description

Daniel found that when we scroll down, the browser sends max_position in query string and the server responses with min_position and has_more_items to tell the browser where to start for the next query.

In this solution, the response from Twitter has "items_html" and we use a library called goquery to parse that.
Note that the "site" of this solution is the API call from Twitter instead of the URL for the page in the last solution.
goquery is basically a search library built on top of Go's net/html package and the CSS Selector library cascadia.

We organized our code into 3 functions and a main function.
1st function (makeConversionRequest):
    makes a single request to the Twitter API. It returns a part of HTML.
2nd function (getConversation): 
    makes all the requests needed to get the whole HTML using the 1st function.    
3rd function (parseHtml):
    parse the "items_html" into tweets. The 1st function makes a single request to the API. The 2nd function makes all the requests needed to get all of the items using the 1st function.

In the main function, we first get the whole conversation by calling getConversation. It gives us all the HTML strings. Then we range over all those HTML strings in the for loop.
Each time through, we call parseHtml to pull out all the tweets from that particular chunk of HTML.
Then we append all those tweets into the bigger tweets slice.

## References

[Errors](https://blog.golang.org/go1.13-errors)
[goquery](https://github.com/PuerkitoBio/goquery)
[robots.txt](https://twitter.com/robots.txt)
[robots.txt validator](https://technicalseo.com/tools/robots-txt/)
[url.Values](https://godoc.org/net/url#Values)
