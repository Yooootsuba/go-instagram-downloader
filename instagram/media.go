package instagram

import (
    "crypto/md5"
    "encoding/hex"
    "github.com/gocolly/colly"
)

type InstagramMedia struct {
    is_video bool
    url      string
}

func NewInstagramMedia(node Node) *InstagramMedia {
    self := new(InstagramMedia)

    self.is_video = node.IsVideo
    if self.is_video == true {
        self.url = node.VideoUrl
    } else {
        self.url = node.DisplayUrl
    }

    return self
}

func (self *InstagramMedia) Download() {
    c := colly.NewCollector()

    c.OnResponse(func(r *colly.Response) {
        sum      := md5.Sum(r.Body)
        filename := hex.EncodeToString(sum[:])

        var extension string
        if self.is_video == true {
            extension = ".mp4"
        } else {
            extension = ".jpg"
        }

        r.Save(filename + extension)
    })

    c.Visit(self.url)
}
