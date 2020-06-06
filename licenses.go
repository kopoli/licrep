// Generated with by licrep version 0.2.1
// https://github.com/kopoli/licrep
// Called with: licrep -o licenses.go --prefix Licrep

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"strings"
)

// LicrepLicense is a representation of an embedded license.
type LicrepLicense struct {
	// Name of the license
	Name string

	// The text of the license
	Text string
}

// LicrepGetLicenses gets a map of LicrepLicenses where the keys are
// the package names.
func LicrepGetLicenses() (map[string]LicrepLicense, error) {
	type EncodedLicense struct {
		Name string
		Text string
	}
	data := map[string]EncodedLicense{

		"github.com/jawher/mow.cli": EncodedLicense{
			Name: "MIT",
			Text: `
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
		"github.com/jawher/mow.cli/internal/container": EncodedLicense{
			Name: "MIT",
			Text: `
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
		"github.com/jawher/mow.cli/internal/flow": EncodedLicense{
			Name: "MIT",
			Text: `
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
		"github.com/jawher/mow.cli/internal/fsm": EncodedLicense{
			Name: "MIT",
			Text: `
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
		"github.com/jawher/mow.cli/internal/lexer": EncodedLicense{
			Name: "MIT",
			Text: `
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
		"github.com/jawher/mow.cli/internal/matcher": EncodedLicense{
			Name: "MIT",
			Text: `
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
		"github.com/jawher/mow.cli/internal/parser": EncodedLicense{
			Name: "MIT",
			Text: `
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
		"github.com/jawher/mow.cli/internal/values": EncodedLicense{
			Name: "MIT",
			Text: `
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
		"github.com/kopoli/appkit": EncodedLicense{
			Name: "MIT",
			Text: `
H4sIAAAAAAAC/1xRzW7jNhC+8yk+5JQAQrrYY2+MRVtEJNKg6HV9pCU6YiuLhkg3yNsXIzu7zZ4Eceb7
HTt4NNKiDp2fksdjI+0TY6t4+ZjD25Dx2D3h+7fv3/DqxtHj1U3/uNkztvXzOaQU4oSQMPjZHz/wNrsp
+77AafYe8YRucPObL5Aj3PSBi59TnBCP2YUpTG9w6OLlg8UT8hASUjzldzd7uKmHSyl2wWXfo4/d9eyn
7DLpncLoEx7z4PHQ3hEPT4tI793IwgSafY7wHvIQrxmzT3kOHXEUCFM3Xnvy8DkewzncFQi+xE8sR1yT
LxafBc6xDyf6+iXW5XocQxoK9IGoj9fsCyR6XNosKMcfcUby48i6eAk+Ycn6y92yQ9YvVGi+V5To5X2I
569JQmKn6zyFNPgF00ekuCj+7btML7R+iuMY3ylaF6c+UKL0J2N0aneM//oly+26U8yhu9W9HODy66r3
URrcOOLo74X5HmGC+1+cmeRTdlMObsQlzove7zGfGbOVQKvXds+NgGyxNfqHLEWJB95Ctg8F9tJWemex
58ZwZQ/Qa3B1wKtUZQHx19aItoU2TDbbWoqygFSreldKtcHLzkJpi1o20ooSVoME71RStETWCLOquLL8
RdbSHgq2llYR51obcGy5sXK1q7nBdme2uhXgqoTSSqq1kWojGqHsM6SC0hA/hLJoK17XJMX4zlbakD+s
9PZg5KayqHRdCtPiRaCW/KUWNyl1wKrmsilQ8oZvxILSthKG0drNHfaVoCfS4wp8ZaVWFGOllTV8ZQtY
bexP6F62ogA3sqVC1kY3BaM69ZpWpCKcEjcWqhpfLqLN8r9rxU9ClILXUm1aAlPEz+Vn9l8AAAD//7MD
VDw4BAAA`,
		},
		"github.com/kopoli/licrep/lib": EncodedLicense{
			Name: "MIT",
			Text: `
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
		"github.com/ryanuber/go-license": EncodedLicense{
			Name: "MIT",
			Text: `
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
		text, err := decode(data[k].Text)
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
