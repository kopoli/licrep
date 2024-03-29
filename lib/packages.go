package licrep

import (
	"errors"
	"go/build"
	"path/filepath"
	"sort"

	"github.com/ryanuber/go-license"
)

const (
	// UnknownLicense denotes the value of License of a Package for which
	// the license is not detected.
	UnknownLicense = "Unknown"
)

// Package reprsents an imported package
type Package struct {
	Name          string
	ImportPath    string
	Dir           string
	License       string
	LicenseString string
}

func findLicense(dir, root string) (ret *license.License, err error) {
	d, _ := filepath.Split(dir)
	d = filepath.Clean(d)
	if d == "" || d == root {
		return nil, errors.New("no license found in parent directories")
	}

	ret, err = license.NewFromDir(dir)
	if err != nil {
		return findLicense(filepath.Join(dir, ".."), root)
	}

	return
}

// GetPackages gets the dependencies of given pkgname in the given directory
func GetPackages(pkgname string, dir string) (ret []Package, err error) {
	imports := map[string]Package{}

	var findImports func(string, string) error

	findImports = func(name, dir string) (err error) {
		pkg, err := build.Import(name, dir, 0)
		if err != nil {
			return err
		}

		if pkg.Goroot {
			return nil
		}

		_, ok := imports[pkg.ImportPath]
		if ok {
			return nil
		}

		p := Package{
			Name:       pkg.Name,
			ImportPath: pkg.ImportPath,
			Dir:        pkg.Dir,
		}

		l, err := findLicense(pkg.Dir, pkg.SrcRoot)
		if err != nil {
			p.License = UnknownLicense
		} else {
			p.License = l.Type
			p.LicenseString = l.Text
		}

		imports[pkg.ImportPath] = p

		for i := range pkg.Imports {
			err = findImports(pkg.Imports[i], pkg.Dir)
		}
		return err
	}

	err = findImports(pkgname, dir)
	if err != nil {
		return nil, err
	}

	var names []string
	for i := range imports {
		names = append(names, i)
	}
	sort.Strings(names)

	ret = make([]Package, len(imports))
	j := 0
	for _, k := range names {
		ret[j] = imports[k]
		j++
	}

	return ret, nil
}
