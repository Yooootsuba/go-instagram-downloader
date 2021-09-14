package instagram

import (
    "encoding/json"
    "github.com/gocolly/colly"
)

type InstagramMediaDownloader struct {
    url string
}

func NewInstagramMediaDownloader(url string) *InstagramMediaDownloader {
    self := new(InstagramMediaDownloader)
    self.url = url
    return self
}

func (self *InstagramMediaDownloader) GetResponse() Response {
    var response Response
    c := colly.NewCollector()

    c.OnResponse(func(r *colly.Response) {
        json.Unmarshal(r.Body, &response)
    })

    c.Visit(self.url + "?__a=1")

    return response
}

func (self *InstagramMediaDownloader) ParseResponse(response Response) []*InstagramMedia {
   var medias []*InstagramMedia

   if response.GraphQl.ShortCodeMedia.EdgeSidecarToChildren.Edges == nil {
       var node Node
       str, _ := json.Marshal(response.GraphQl.ShortCodeMedia)
       json.Unmarshal(str, &node)
       medias = append(medias, NewInstagramMedia(node))
       return medias
   }

   for _, edge := range response.GraphQl.ShortCodeMedia.EdgeSidecarToChildren.Edges {
       medias = append(medias, NewInstagramMedia(edge.Node))
   }

   return medias
}

func (self *InstagramMediaDownloader) Download() {
    response := self.GetResponse()
    medias   := self.ParseResponse(response)

    for _, media := range medias {
        media.Download()
    }
}
