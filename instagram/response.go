package instagram

type Node struct {
    IsVideo    bool   `json:"is_video"`
    VideoUrl   string `json:"video_url"`
    DisplayUrl string `json:"display_url"`
}

type Edges struct {
    Node Node `json:"node"`
}

type EdgeSidecarToChildren struct {
    Edges []Edges `json:"edges"`
}

type ShortCodeMedia struct {
    IsVideo               bool                  `json:"is_video"`
    VideoUrl              string                `json:"video_url"`
    DisplayUrl            string                `json:"display_url"`
    EdgeSidecarToChildren EdgeSidecarToChildren `json:"edge_sidecar_to_children"`
}

type GraphQl struct {
    ShortCodeMedia ShortCodeMedia `json:"shortcode_media"`
}

type Response struct {
    GraphQl GraphQl `json:"graphql"`
}
