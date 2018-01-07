// Generated with by licrep version 0.1.0
// https://github.com/kopoli/licrep
// Called with: licrep -o licenses.go -p main --prefix Licrep

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"strings"
)

type LicrepLicense struct {
	// Name of the license
	Name string

	// The text of the license
	Text string
}

func LicrepGetLicenses() (map[string]LicrepLicense, error) {
	type Data struct {
		Name string
		Data string
	}
	data := map[string]Data{

		"github.com/davecgh/go-spew/spew": Data{
			Name: "ISC",
			Data: `
H4sIAAAAAAAC/1SQz27bOBjE73yKQU67gNab5NBTUZSRPlkEaNIlqRg5KhYdE3BEQaId+O0LKk2bXvgH
/GbmxxG2hAx7P8yesTKO1ym8HBP+2f+L+9u7+//ub+++oOouHmU8ncIw42vfXfz3vOxXQ/fqvzG29dNr
mOcQB6SI8+wL7ON4LfAa+3C4FuiG/v84oQ9zmsLzOXmkY5gxx0N66yaPQ5zQDVc2nqcxzh5vIR0Rp2WP
54SD9wgzjn7yz1e8TN2QfF9gnOIl9L5HOnYJ6ejRPceLZ/vf/xhiCnuf898Txz+kH0/j6LsJYUB3OmXs
4OcVY64hWF27HTcEYbE1+lFUVOGGWwh7A64q5CHeukYbVMKWkouNBZcSO24MV06QZTvhGhhac1PBabhG
2E/GqpRtJdR6UYnNVgqqPqmha7YhUzZcOf4gpHBPS3AtnCJrVxAKSoMeSTnYJpt8YnogSMEfJKHWhnH1
BLulUnBZoBKGSldAqI+TNii1svSjJeUEl6j4hq8zgkGW/rqyXcOd1fRIBoZsK12mr43eQGqbgdFaKlBx
x7N0a3QtnC2wa8g1ZDIxV4yXTmiVp0utnOGZQNFaijWpkrJQL9NOGyd0a/EuKMCNsDlRty6XoxfDUitF
745L3bmD1i42WzK1Nhu+uNZ/179iPwMAAP//T6K2Jf4CAAA=`,
		},
		"github.com/jawher/mow.cli": Data{
			Name: "MIT",
			Data: `
H4sIAAAAAAAC/3xRvY7jNhjs+RSDre4AYvODVOm4Em0xkUiDos9xSUv0ioEsGiIVw28fUPbeYVNEjcDv
Z+abGfz/ZwaHRhjUvnNTdIQU4Xqf/fuQ8KX7il9//uU3ij/sbXAzmrDEaAnZufniY/Rhgo8Y3OxOd7zP
dkqupzjPziGc0Q12fncUKcBOd1zdHMOEcErWT356h0UXrncSzkiDj4jhnG52drBTDxtj6LxNrkcfuuXi
pmRT5jv70UV8SYPDS/vcePm6kvTOjsRPyL2PFm4+DWFJmF1Ms+8yBoWfunHp8w0f7dFf/JMhr6/6I0kB
S3R0vZPiEnp/zn+3yroup9HHgaL3Gfq0JEcRc3E1kmYdP4UZ0Y0j6cLVu4hV64/r1pl8+jUbmp4WxVy5
DeHyWYmP5LzMk4+DW3f6gBhWxr9dl3Ilj5/DOIZbltaFqfdZUfydkJyyPYV/3KrlEe8Uku8edq8BXH+k
+mzFwY4jTu5pmOvhJ5JLH3LmTB+TnZK3I65hXvn+K/OVEFNxtGpjDkxziBY7rb6Jkpd4YS1E+0JxEKZS
e4MD05pJc4TagMkj/hSypOB/7TRvWyhNRLOrBS8phCzqfSnkFm97A6kMatEIw0sYhUz4hBK8zWAN10XF
pGFvohbmSMlGGJkxN0qDYce0EcW+Zhq7vd6ploPJElJJITdayC1vuDSvEBJSgX/j0qCtWF1nKsL2plI6
34dC7Y5abCuDStUl1y3eOGrB3mr+oJJHFDUTDUXJGrbl65YyFdckjz2uw6HiuZT5mAQrjFAyyyiUNJoV
hsIobb6vHkTLKZgWbTZko1VDSbZTbfKIkHlP8gdKthqfElF6fe9b/h0QJWe1kNsWQn6K75WQfwMAAP//
eVpV+1MEAAA=`,
		},
		"github.com/kopoli/go-util": Data{
			Name: "MIT",
			Data: `
H4sIAAAAAAAC/1xRzW7jNhC+8yk+5JQAQvpz6KE3xqItIhJpUPS6PtISHbGVRUOkG+Tti5Gd3WZPgjjz
/Y4dPBppUYfOT8njsZH2ibFVvHzM4W3IeOye8Puvv/2BVzeOHq9u+sfNnrGtn88hpRAnhITBz/74gbfZ
Tdn3BU6z94gndIOb33yBHOGmD1z8nOKEeMwuTGF6g0MXLx8snpCHkJDiKb+72cNNPVxKsQsu+x597K5n
P2WXSe8URp/wmAePh/aOeHhaRHrvRhYm0OxzhPeQh3jNmH3Kc+iIo0CYuvHak4fP8RjO4a5A8CV+Yjni
mnyx+Cxwjn040dcvsS7X4xjSUKAPRH28Zl8g0ePSZkE5fokzkh9H1sVL8AlL1h/ulh2yfqFC872iRC/v
Qzx/TRISO13nKaTBL5g+IsVF8W/fZXqh9VMcx/hO0bo49YESpT8Zo1O7Y/zXL1lu151iDt2t7uUAlx9X
vY/S4MYRR38vzPcIE9z/4swkn7KbcnAjLnFe9H6O+cyYrQRavbZ7bgRki63R32QpSjzwFrJ9KLCXttI7
iz03hit7gF6DqwNepSoLiL+2RrQttGGy2dZSlAWkWtW7UqoNXnYWSlvUspFWlLAaJHinkqIlskaYVcWV
5S+ylvZQsLW0ijjX2oBjy42Vq13NDbY7s9WtAFcllFZSrY1UG9EIZZ8hFZSG+CaURVvxuiYpxne20ob8
YaW3ByM3lUWl61KYFi8CteQvtbhJqQNWNZdNgZI3fCMWlLaVMIzWbu6wrwQ9kR5X4CsrtaIYK62s4Stb
wGpjv0P3shUFuJEtFbI2uikY1anXtCIV4ZS4sVDV+HIRbZb/XSu+E6IUvJZq0xKYIn4uP/8XAAD//6vn
VXQ3BAAA`,
		},
		"github.com/kopoli/licrep": Data{
			Name: "MIT",
			Data: `
H4sIAAAAAAAC/1xRzW6rOBTeI/EOn7pqJdT52Yw0Ozc4wSrYkXFuJksHnOIZgiPsTNW3Hx2S3ju9K4TP
+X6PGRwaYVD7zk3R4bER5inP8mwVLh+zfxsSHrsn/P7rb3/g1Y6jw6ud/rGzo52tm88+Rh8m+IjBze74
gbfZTsn1BU6zcwgndIOd31yBFGCnD1zcHMOEcEzWT356g0UXLh95Fk5Ig4+I4ZTe7exgpx42xtB5m1yP
PnTXs5uSTSR48qOLeEyDw0N7Rzw8LSq9s2Oe+Qk0/Jzh3achXBNmF9PsOyIp4KduvPbk4nM8+rO/SxB8
qSDmWQq4RlcsVgucQ+9P9HVLssv1OPo4FOg9cR+vyRWI9LiUWlCSX8KM6MYxz7pw8S5iifvD37JE7i9U
arrXFOnlfQjnr1l8zLPTdZ58HNwC6gNiWDT/dl2iF9o/hXEM75SuC1PvKVT8kw5HV7fH8K9b8tzOPIXk
u1vpyxkuP457H8XBjiOO7t6a6+En2P9HmslBTHZK3o64hHmR/Dnq82Kh4mjV2uyZ5hAttlp9EyUv8cBa
iPahwF6YSu0M9kxrJs0Bag0mD3gVsizA/9pq3rZQOs9Es60FLwsIuap3pZAbvOwMpDKoRSMML2EUSPHO
JXhLbA3Xq4pJw15ELcyhyLO1MJJY10qDYcu0EatdzTS2O71VLQeTJaSSQq61kBvecGmeISSkAv/GpUFb
sbomrTxjO1MpTRaxUtuDFpvKoFJ1yXWLF45asJea37TkAauaiaZAyRq24QtKmYrrPKO9m0HsK05vpMgk
2MoIJSnJSkmj2coUMEqb79i9aHkBpkVLnay1aoo8o07VmnaEJKDkNxrqG1/OovTyv2v5d0aUnNVCbloC
Lyk/t5/z7L8AAAD//1HmlTNNBAAA`,
		},
		"github.com/kopoli/licrep/lib": Data{
			Name: "MIT",
			Data: `
H4sIAAAAAAAC/1xRzW6rOBTeI/EOn7pqJdT52Yw0Ozc4wSrYkXFuJksHnOIZgiPsTNW3Hx2S3ju9K4TP
+X6PGRwaYVD7zk3R4bER5inP8mwVLh+zfxsSHrsn/P7rb3/g1Y6jw6ud/rGzo52tm88+Rh8m+IjBze74
gbfZTsn1BU6zcwgndIOd31yBFGCnD1zcHMOEcEzWT356g0UXLh95Fk5Ig4+I4ZTe7exgpx42xtB5m1yP
PnTXs5uSTSR48qOLeEyDw0N7Rzw8LSq9s2Oe+Qk0/Jzh3achXBNmF9PsOyIp4KduvPbk4nM8+rO/SxB8
qSDmWQq4RlcsVgucQ+9P9HVLssv1OPo4FOg9cR+vyRWI9LiUWlCSX8KM6MYxz7pw8S5iifvD37JE7i9U
arrXFOnlfQjnr1l8zLPTdZ58HNwC6gNiWDT/dl2iF9o/hXEM75SuC1PvKVT8kw5HV7fH8K9b8tzOPIXk
u1vpyxkuP457H8XBjiOO7t6a6+En2P9HmslBTHZK3o64hHmR/Dnq82Kh4mjV2uyZ5hAttlp9EyUv8cBa
iPahwF6YSu0M9kxrJs0Bag0mD3gVsizA/9pq3rZQOs9Es60FLwsIuap3pZAbvOwMpDKoRSMML2EUSPHO
JXhLbA3Xq4pJw15ELcyhyLO1MJJY10qDYcu0EatdzTS2O71VLQeTJaSSQq61kBvecGmeISSkAv/GpUFb
sbomrTxjO1MpTRaxUtuDFpvKoFJ1yXWLF45asJea37TkAauaiaZAyRq24QtKmYrrPKO9m0HsK05vpMgk
2MoIJSnJSkmj2coUMEqb79i9aHkBpkVLnay1aoo8o07VmnaEJKDkNxrqG1/OovTyv2v5d0aUnNVCbloC
Lyk/t5/z7L8AAAD//1HmlTNNBAAA`,
		},
		"github.com/pkg/errors": Data{
			Name: "FreeBSD",
			Data: `
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
		"github.com/ryanuber/go-license": Data{
			Name: "MIT",
			Data: `
H4sIAAAAAAAC/1xRzW7jNhC+8yk+5JQAQvqDnnpjLNoiKpEGRa/royzREQuZNES6gd++GNnZbfYkiDPf
79jRoZEWte9dSA7PjbQvjK3i5Tb79zHjuX/B77/+9gfMrQvYHd3M2NbNZ5+SjwE+YXSzO97wPnchu6HA
aXYO8YR+7OZ3VyBHdOGGi5tTDIjH3Pngwzs69PFyY/GEPPqEFE/5o5sdujCgSyn2vstuwBD769mF3GXS
O/nJJTzn0eGpfSCeXhaRwXUT8wE0+xzhw+cxXjNml/Lse+Io4EM/XQfy8Dme/Nk/FAi+RE8sR1yTKxaf
Bc5x8Cf6uiXW5XqcfBoLDJ6oj9fsCiR6XJosKMcvcUZy08T6ePEuYcn6w92yQ9YvVGh+VJTo5WOM569J
fGKn6xx8Gt2CGSJSXBT/cX2mF1o/xWmKHxStj2HwlCj9yRiduTvGf92S5X7ZELPv73UvB7j8uOpjlMZu
mnB0j8LcAB/Q/S/OTPIpdyH7bsIlzovezzFfGbOVQKvXds+NgGyxNfqbLEWJJ95Ctk8F9tJWemex58Zw
ZQ/Qa3B1wF9SlQXE31sj2hbaMNlsaynKAlKt6l0p1QZvOwulLWrZSCtKWA0SfFBJ0RJZI8yq4sryN1lL
eyjYWlpFnGttwLHlxsrVruYG253Z6laAqxJKK6nWRqqNaISyr5AKSkN8E8qirXhdkxTjO1tpQ/6w0tuD
kZvKotJ1KUyLN4Fa8rda3KXUAauay6ZAyRu+EQtK20oYRmt3d9hXgp5IjyvwlZVaUYyVVtbwlS1gtbHf
oXvZigLcyJYKWRvdFIzq1GtakYpwStxZqGp8uYg2y/+uFd8JUQpeS7VpCUwRP5df2X8BAAD//9UdrcM0
BAAA`,
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

	ret := make(map[string]LicrepLicense)

	for k := range data {
		text, err := decode(data[k].Data)
		if err != nil {
			return nil, err
		}
		ret[k] = LicrepLicense{
			Name: data[k].Name,
			Text: text,
		}
	}

	return ret, nil
}
