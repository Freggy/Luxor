package generator

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"math/rand"
)

type Noun struct {
	// Article of the word. Can be either DER, DIE or DAS
	Article string `json:"article"`
	// The actual word.
	Word string `json:"noun"`
}

type NounList []Noun

type AdjectiveList []string

var (
	nouns      NounList
	adjectives AdjectiveList
)

func init() {
	marshallAssets(&nouns, &adjectives)
}

// GenerateName generates a unique name for the given input string.
func GenerateName(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	hash.Sum(nil)
	seed := binary.BigEndian.Uint64(hash.Sum(nil))
	random := rand.New(rand.NewSource(int64(seed)))
	noun := nouns[random.Intn(len(nouns))]
	adjective := adjectives[random.Intn(len(adjectives))]
	return noun.Article + "_" + adjective + "_" + noun.Word
}

func marshallAssets(nouns *NounList, adjectives *AdjectiveList) {
	_ = json.Unmarshal([]byte(Nouns), &nouns)
	_ = json.Unmarshal([]byte(Adjectives), &adjectives)
}
