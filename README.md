# swear-free
This is a go library to censor all the swear words in a given text.

# Installation

```
go get github.com/ohhhthatvarun/swear-free
```

# Languages/Locale supported
- English (En)

# Usage

* To get censored version of the text.

```
censoredText = swearFree.CensorWord("Your bitch ass text here")
fmt.Println(censoredText) // Your ***** *** text here
```

* To change the censor character

```
swearFree.SetReplaceCharacter("#")
censoredText = swearFree.CensorWord("Your bitch ass text here")
fmt.Println(censoredText) // Your ##### ### text here
```

* To change the locale (Future)

```
swearFree.SetLocale("Es")
```
