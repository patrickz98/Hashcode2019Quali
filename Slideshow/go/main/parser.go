package main

import (
	"../simple"
	"github.com/golang-collections/collections/set"
	"io/ioutil"
	"strconv"
	"strings"
)

func Init(path string) SlideShow {
	dat, err := ioutil.ReadFile(path)
	simple.CheckErr(err)

	lines := strings.SplitAfter(string(dat), "\n")
	head, lines := lines[0], lines[1:]

	items, err := strconv.Atoi(strings.TrimSuffix(head, "\n"))
	simple.CheckErr(err)

	photos := make([]*Photo, items)

	for inx, line := range lines {

		line = strings.TrimSuffix(line, "\n")
		parts := strings.Split(line, " ")

		if len(parts) < 3 {
			continue
		}

		tagsLen, err := strconv.Atoi(parts[ 1 ])
		simple.CheckErr(err)

		tags := set.New()

		for _, tag := range parts[2:] {
			tags.Insert(tag)
		}

		photo := &Photo{
			ID: inx,
			Orientation: parts[ 0 ],
			Tags: tags,
			TagsLen: tagsLen,
		}

		photos[ inx ] = photo
	}

	slideshow := SlideShow{
		Photos: photos,
	}

	return slideshow
}
