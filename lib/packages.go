package licrep

import (
	"fmt"
	"go/build"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/ryanuber/go-license"
)

type Package struct {
	Name    string
	Dir     string
	License string
	LicenseString string
}

func findLicense(dir string) (ret *license.License, err error) {
	pkg, err := build.Import(".", dir, 0)
	if err != nil {
		return
	}

	ret, err = license.NewFromDir(pkg.Dir)
	if err != nil {
		return findLicense(filepath.Join(pkg.Dir, ".."))
	}

	return
}

func GetPackages(pkgname string, dir string) (ret []Package, err error) {
	imports := map[string]Package{}

	var findImports func(string, string) error

	findImports = func(name, dir string) (err error) {
		pkg, err := build.Import(name, dir, 0)
		if err != nil {
			return
		}

		if pkg.Goroot {
			return
		}

		fmt.Println(spew.Sdump(pkg))

		_, ok := imports[pkg.Name]
		if ok {
			return
		}

		p := Package{
			Name:    pkg.Name,
			Dir:     pkg.Dir,
		}

		l, err := findLicense(pkg.Dir)
		if err != nil {
			p.License = "Unknown"
		} else {
			p.License = l.Type
			p.LicenseString = l.Text
		}

		imports[pkg.Name] = p

		for i := range pkg.Imports {
			err = findImports(pkg.Imports[i], pkg.Dir)
			// if err != nil {

			// }
		}
		return
	}

	err = findImports(pkgname, dir)

	fmt.Println(spew.Sdump(imports))
	return
}
