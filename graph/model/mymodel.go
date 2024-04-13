package model

import (
	"fmt"
	"io"
	"net/url"
)

type MyURL struct {
	url.URL
}

func (u MyURL) MarshalGQL(w io.Writer) {
	io.WriteString(w, fmt.Sprintf(`"%s"`, u.URL.String()))
}
