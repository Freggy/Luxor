package app

import (
	"fmt"
	"github.com/kalafut/imohash"
	"os"
	"path/filepath"
)

type FileModel struct {
	Path  string      `json:"path"`
	Hash  string      `json:"hash,omitempty"`
	Other []FileModel `json:"other,omitempty"`
}

func BuildModel(service *Service) ([]FileModel, error) {
	m, err := buildDef(service.Files)

	if err != nil {
		return nil, err
	}
	return m, nil
}

func buildDef(defs []FileDef) ([]FileModel, error) {
	models := make([]FileModel, 0)

	for _, def := range defs {
		result, err := os.Stat(def.Path)

		if err != nil {
			return nil, err
		}

		if result.IsDir() {
			if def.Files == nil {
				err := filepath.Walk(def.Path, func(path string, info os.FileInfo, err error) error {
					if path != def.Path {
						def.Files = append(def.Files, FileDef{Path: path})
						if info.IsDir() {
							return filepath.SkipDir
						}
					}
					return nil
				})

				if err != nil {
					return nil, err
				}
			}

			other, err := buildDef(def.Files)

			if err != nil {
				return nil, err
			}

			m := FileModel{
				Path:  def.Path,
				Other: other,
			}

			models = append(models, m)

		} else {
			hash, err := imohash.SumFile(def.Path)

			if err != nil {
				return nil, err
			}

			m := FileModel{
				Path: def.Path,
				Hash: fmt.Sprintf("%016x", hash),
			}

			models = append(models, m)
		}
	}
	return models, nil
}
