package parquet

import (
	"encoding/binary"
	"os"
)

type ParquetWriter struct {
	file      *os.File
	schema    []*SchemaElement
	rowGroups []*RowGroup
}

func NewParquetWriter(filename string, schema []*SchemaElement) (*ParquetWriter, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	//write header
	header := [4]byte{'P', 'A', 'R', '1'}
	if _, err := file.Write(header[:]); err != nil {
		return nil, err
	}

	return &ParquetWriter{
		file:   file,
		schema: schema,
	}, nil
}

func (w *ParquetWriter) WriteRowGroup(data interface{}) error {
	rg := &RowGroup{
		Columns: make([]*ColumnChunk, len(w.schema)),
	}

	for i, se := range w.schema {
		rg.Columns[i] = &ColumnChunk{
			MetaData: &ColumnMetaData{
				Type:         se.Type,
				PathInSchema: []string{se.Name},
				NumValues:    10,
			},
			Data: make([]byte, 100),
		}
	}

	w.rowGroups = append(w.rowGroups, rg)
	return nil
}

func (w *ParquetWriter) Close() error {
	// Napiši footer
	footer := &FileFooter{
		RowGroups: make([]*RowGroupMeta, len(w.rowGroups)),
		Version:   1,
		Metadata: &FileMetaData{
			Version: 1,
			Schema:  w.schema,
			NumRows: int64(len(w.rowGroups)) * 10, // 10 redova po grupi
		},
	}

	// Pretvori footer u bytes
	footerBytes := serializeFooter(footer)

	// Napiši footer
	if _, err := w.file.Write(footerBytes); err != nil {
		return err
	}

	footerLen := make([]byte, 4)
	binary.LittleEndian.PutUint32(footerLen, uint32(len(footerBytes)))
	if _, err := w.file.Write(footerLen); err != nil {
		return err
	}

	magic := [4]byte{'P', 'A', 'R', '1'}
	if _, err := w.file.Write(magic[:]); err != nil {
		return err
	}

	return w.file.Close()
}

func serializeFooter(footer *FileFooter) []byte {
}
