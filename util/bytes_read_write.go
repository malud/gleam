package util

import (
	"encoding/binary"
	"fmt"
	"io"
)

func ReadMessage(reader io.Reader) (m []byte, err error) {
	var length int32
	err = binary.Read(reader, binary.LittleEndian, &length)
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Errorf("Failed to read message length: %v", err)
		return
	}
	if length == -1 {
		return nil, io.EOF
	}
	if length == 0 {
		return
	}
	m = make([]byte, length)
	_, err = io.ReadFull(reader, m)
	if err == io.EOF {
		return
	}
	if err != nil {
		return nil, fmt.Errorf("Failed to read message content: %v", err)
	}
	return m, nil
}

func WriteEOFMessage(writer io.Writer) (err error) {
	if err = binary.Write(writer, binary.LittleEndian, int32(-1)); err != nil {
		return fmt.Errorf("Failed to write message length: %v", err)
	}
	return
}

func WriteMessage(writer io.Writer, m []byte) (err error) {
	if err = binary.Write(writer, binary.LittleEndian, int32(len(m))); err != nil {
		return fmt.Errorf("Failed to write message length: %v", err)
	}
	if _, err = writer.Write(m); err != nil {
		return fmt.Errorf("Failed to write message content: %v", err)
	}
	return
}