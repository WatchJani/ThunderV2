package parquet

type ParquetFile struct {
	Header    *FileHeader
	RowGroups []*RowGroup
	Footer    *FileFooter
}

type FileHeader struct {
	Magic [4]byte
}

type RowGroup struct {
	Columns []*ColumnChunk
}

type ColumnChunk struct {
	MetaData *ColumnMetaData
	Data     []byte
}

type ColumnMetaData struct {
	Type         int32
	Encodings    []int32
	PathInSchema []string
	NumValues    int64
}

type FileFooter struct {
	RowGroups []*RowGroupMeta
	Version   int32
	Metadata  *FileMetaData
}

type RowGroupMeta struct {
	Columns []*ColumnMetaData
	NumRows int64
}

type FileMetaData struct {
	Version int32
	Schema  []*SchemaElement
	NumRows int64
}

type SchemaElement struct {
	Name string
	Type int32
}
