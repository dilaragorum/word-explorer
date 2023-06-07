package vocabulary

type SearchArgs struct {
	Page       int    `json:"page"`
	Size       int    `json:"size"`
	SearchWord string `json:"search_word"`
}

func (args *SearchArgs) SetDefaults() {
	if args.Page <= 0 {
		args.Page = 1
	}
	if args.Size <= 0 {
		args.Size = 20
	}
}
