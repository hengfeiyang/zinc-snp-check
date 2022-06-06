module github.com/hengfeiyang/zinc-snp-check

go 1.17

require (
	github.com/blugelabs/bluge v0.1.9
	github.com/stretchr/testify v1.7.1 // indirect
)

require (
	github.com/RoaringBitmap/roaring v0.9.4 // indirect
	github.com/bits-and-blooms/bitset v1.2.0 // indirect
	github.com/blevesearch/mmap-go v1.0.4 // indirect
	github.com/blevesearch/vellum v1.0.7 // indirect
	github.com/blugelabs/bluge_segment_api v0.2.0 // indirect
	github.com/blugelabs/ice v1.0.0 // indirect
	github.com/klauspost/compress v1.15.2 // indirect
	github.com/mschoch/smat v0.2.0 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	gopkg.in/yaml.v3 v3.0.0 // indirect
)

replace github.com/blugelabs/bluge => github.com/zinclabs/bluge v1.1.3

replace github.com/blugelabs/ice => github.com/zinclabs/ice v1.1.1

replace github.com/blugelabs/bluge_segment_api => github.com/zinclabs/bluge_segment_api v1.0.0
