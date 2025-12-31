package extraction

// Document holds PDF metadata
type Document struct {
	Title     string
	Author    string
	Creator   string
	PageCount int
	FilePath  string
	FileHash  string
	FileSize  int
	Language  string
}

// Page holds single page data
type Page struct {
	PageNumber int
	RawText    string
	Width      float64
	Height     float64
}

// BlockType represent content type
type BlockType string

const (
	BlockParagraph BlockType = "paragraph"
	BlockHeading   BlockType = "heading"
	BlockImage     BlockType = "image"
	BlockTable     BlockType = "table"
	BlockList      BlockType = "list"
)

// ContentBlock is a piece of content of a page
type ContentBlock struct {
	PageNumber int
	Type       BlockType
	Content    string
	Sequence   int
	X, Y       float64
	Width      float64
	Height     float64
	Fontsize   int
	FontName   string
	Confidence int
}

// OutlineItem represent TOC entry
type OutlineItem struct {
	Title      string
	Level      int
	PageNumber int
	Children   []OutlineItem
}

// Image holds extracted image info
type Image struct {
	PageNumber int
	FilePath   string
	Format     string
	Width      int
	Height     int
}

// Extraction ,final output
type ExtractionResult struct {
	Document     Document
	Page         []Page
	ContentBlock []ContentBlock
	Outline      []OutlineItem
	Image        []Image
}
