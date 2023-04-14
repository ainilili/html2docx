package html2docx

type Option func(data map[string]interface{})

func WithWidth(width int) Option {
	return func(data map[string]interface{}) {
		data["width"] = width
	}
}

func WithHeight(height int) Option {
	return func(data map[string]interface{}) {
		data["height"] = height
	}
}

func WithOrient(orient string) Option {
	return func(data map[string]interface{}) {
		data["orient"] = orient
	}
}

func WithMarginsTop(top int) Option {
	return func(data map[string]interface{}) {
		data["margins"].(map[string]interface{})["top"] = top
	}
}

func WithMarginsRight(right int) Option {
	return func(data map[string]interface{}) {
		data["margins"].(map[string]interface{})["right"] = right
	}
}

func WithMarginsBottom(bottom int) Option {
	return func(data map[string]interface{}) {
		data["margins"].(map[string]interface{})["bottom"] = bottom
	}
}

func WithMarginsLeft(left int) Option {
	return func(data map[string]interface{}) {
		data["margins"].(map[string]interface{})["left"] = left
	}
}

func WithMarginsHeader(header int) Option {
	return func(data map[string]interface{}) {
		data["margins"].(map[string]interface{})["header"] = header
	}
}

func WithMarginsFooter(footer int) Option {
	return func(data map[string]interface{}) {
		data["margins"].(map[string]interface{})["footer"] = footer
	}
}

func WithMarginsGutter(gutter int) Option {
	return func(data map[string]interface{}) {
		data["margins"].(map[string]interface{})["gutter"] = gutter
	}
}
