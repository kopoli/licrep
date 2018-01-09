// Generated with by licrep version 0.1.0
// https://github.com/kopoli/licrep
// Called with: licrep -o licenses.go

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"strings"
)

type License struct {
	// Name of the license
	Name string

	// The text of the license
	Text string
}

func GetLicenses() (map[string]License, error) {
	type EncodedLicense struct {
		Name string
		Text string
	}
	data := map[string]EncodedLicense{

		"github.com/pkg/errors": EncodedLicense{
			Name: "FreeBSD",
			Text: `
H4sIAAAAAAAC/5SSQY/bNhOG7/wVL3LKfhD8tQV6aYqitDS2BpBJlaTW8VFrcWMCtrSQ6A323xdkdhGn
DVr05LHEmXnehyqnp5c5fDpFvD/e4acffvy5QNU/e5QnP/oX/Dr0z/73Y/6zGn38TcjzGbljwewXPz/7
YSWE8UNY4hwerjFMI/pxwHXxCCOW6ToffX7yEMZ+fsHjNF+WAp9DPGGa8+90jeIyDeExHPs0oEA/ezz5
+RJi9AOe5uk5DH5APPUR8eTxOJ3P0+cwfsJxGoeQmpbcdPHxFyH+h2+JFkyPbyjHafC4XJeI2cc+jHle
/zA9p1dvOsYphqMvEE9hEcA5LDGNuF02Dn8hGcJyPPfh4ufV9wjCeGvgjeBpnobr0f8ThEDG+K8QeI02
TMfrxY8xmxVIPf+fZkzx5Gdc+ujn0J+Xr47zxeTGG/yVEK5mC6s3bi8NgS1ao++5ogrrA1xNKHV7MLyt
HWrdVGQspKpQauUMrzunjcU7acH2nUgvpDqAPraGrIU24F3bMFXYS2Okcky2AKuy6SpW2wLrzkFph4Z3
7KiC00VaKv7eBr3BjkxZS+Xkmht2hwyyYafSro02kGilcVx2jTRoO9NqS5CGRMW2bCTvqFqBFZQG3ZNy
sLVsmu+mTOzfZFwTGpbrhkTepA6o2FDpUpyvVckVKSebAralklNBH2nXNtIciteZlv7oSDmWjajkTm7J
4v2/KGmNLjtDu8SsN7Dd2jp2nSNsta6SaGHJ3HNJ9gMabbOtzlKBSjqZF7dGb9jZD6led5azNFaOjOla
x1rdodZ7uicjStlZqrJdrXJUV5M2hzQ0OcjyC+xrcjWZJDSbkkmBdYZLd3NMaAOnjbvJCEXbhrekSko0
Ok3Zs6U7SMM2HeAva/fyAN3lyOmOOksilzdfbJFvEryBrO45Yb8ebrW1/PqdZGVljS+6V+LPAAAA//+5
j2aZIAUAAA==`,
		},
	}

	decode := func(input string) (string, error) {
		data := &bytes.Buffer{}
		br := base64.NewDecoder(base64.StdEncoding, strings.NewReader(input))

		r, err := gzip.NewReader(br)
		if err != nil {
			return "", err
		}

		_, err = io.Copy(data, r)
		if err != nil {
			return "", err
		}

		// Make sure the gzip is decoded successfully
		err = r.Close()
		if err != nil {
			return "", err
		}
		return data.String(), nil
	}

	ret := make(map[string]License)

	for k := range data {
		text, err := decode(data[k].Text)
		if err != nil {
			return nil, err
		}
		ret[k] = License{
			Name: data[k].Name,
			Text: text,
		}
	}

	return ret, nil
}
