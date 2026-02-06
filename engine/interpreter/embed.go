package interpreter

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

const (
	EMBED_MAGIC   = "KIDLANG_EMBEDDED_SCRIPT"
	EMBED_VERSION = uint32(1)
)

type EmbeddedHeader struct {
	Magic       [24]byte
	Version     uint32
	ScriptSize  uint32
	Checksum    [32]byte
	NameLength  uint32
	NameBytes   []byte
	ScriptBytes []byte
}

func BuildEmbeddedBinary(scriptPath, outputPath string) error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get current executable path: %w", err)
	}

	scriptData, err := os.ReadFile(scriptPath)
	if err != nil {
		return fmt.Errorf("failed to read script file: %w", err)
	}

	exeData, err := os.ReadFile(exePath)
	if err != nil {
		return fmt.Errorf("failed to read executable: %w", err)
	}

	scriptName := filepath.Base(scriptPath)
	nameBytes := []byte(scriptName)

	checksum := sha256.Sum256(scriptData)

	header := EmbeddedHeader{
		Version:     EMBED_VERSION,
		ScriptSize:  uint32(len(scriptData)),
		Checksum:    checksum,
		NameLength:  uint32(len(nameBytes)),
		NameBytes:   nameBytes,
		ScriptBytes: scriptData,
	}
	copy(header.Magic[:], EMBED_MAGIC)

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	if _, err := outFile.Write(exeData); err != nil {
		return fmt.Errorf("failed to write executable data: %w", err)
	}

	if err := binary.Write(outFile, binary.LittleEndian, header.Magic); err != nil {
		return fmt.Errorf("failed to write magic: %w", err)
	}
	if err := binary.Write(outFile, binary.LittleEndian, header.Version); err != nil {
		return fmt.Errorf("failed to write version: %w", err)
	}
	if err := binary.Write(outFile, binary.LittleEndian, header.ScriptSize); err != nil {
		return fmt.Errorf("failed to write script size: %w", err)
	}
	if err := binary.Write(outFile, binary.LittleEndian, header.Checksum); err != nil {
		return fmt.Errorf("failed to write checksum: %w", err)
	}
	if err := binary.Write(outFile, binary.LittleEndian, header.NameLength); err != nil {
		return fmt.Errorf("failed to write name length: %w", err)
	}
	if _, err := outFile.Write(nameBytes); err != nil {
		return fmt.Errorf("failed to write name: %w", err)
	}
	if _, err := outFile.Write(scriptData); err != nil {
		return fmt.Errorf("failed to write script: %w", err)
	}

	if err := outFile.Chmod(0755); err != nil {
		return fmt.Errorf("failed to set executable permissions: %w", err)
	}

	return nil
}

func ExtractEmbeddedScript() (string, []byte, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", nil, fmt.Errorf("failed to get executable path: %w", err)
	}

	exeFile, err := os.Open(exePath)
	if err != nil {
		return "", nil, fmt.Errorf("failed to open executable: %w", err)
	}
	defer exeFile.Close()

	stat, err := exeFile.Stat()
	if err != nil {
		return "", nil, fmt.Errorf("failed to stat executable: %w", err)
	}
	fileSize := stat.Size()

	searchStart := fileSize - (1024 * 1024)
	if searchStart < 0 {
		searchStart = 0
	}

	if _, err := exeFile.Seek(searchStart, 0); err != nil {
		return "", nil, fmt.Errorf("failed to seek: %w", err)
	}

	tailData, err := io.ReadAll(exeFile)
	if err != nil {
		return "", nil, fmt.Errorf("failed to read tail: %w", err)
	}

	magicBytes := []byte(EMBED_MAGIC)
	magicPos := bytes.LastIndex(tailData, magicBytes)
	if magicPos == -1 {
		return "", nil, fmt.Errorf("no embedded script found")
	}

	headerStart := searchStart + int64(magicPos)
	if _, err := exeFile.Seek(headerStart, 0); err != nil {
		return "", nil, fmt.Errorf("failed to seek to header: %w", err)
	}

	var header EmbeddedHeader
	if err := binary.Read(exeFile, binary.LittleEndian, &header.Magic); err != nil {
		return "", nil, fmt.Errorf("failed to read magic: %w", err)
	}
	if err := binary.Read(exeFile, binary.LittleEndian, &header.Version); err != nil {
		return "", nil, fmt.Errorf("failed to read version: %w", err)
	}
	if err := binary.Read(exeFile, binary.LittleEndian, &header.ScriptSize); err != nil {
		return "", nil, fmt.Errorf("failed to read script size: %w", err)
	}
	if err := binary.Read(exeFile, binary.LittleEndian, &header.Checksum); err != nil {
		return "", nil, fmt.Errorf("failed to read checksum: %w", err)
	}
	if err := binary.Read(exeFile, binary.LittleEndian, &header.NameLength); err != nil {
		return "", nil, fmt.Errorf("failed to read name length: %w", err)
	}

	if header.NameLength > 1024 {
		return "", nil, fmt.Errorf("name length too large: %d", header.NameLength)
	}
	if header.ScriptSize > 100*1024*1024 {
		return "", nil, fmt.Errorf("script size too large: %d", header.ScriptSize)
	}

	header.NameBytes = make([]byte, header.NameLength)
	if _, err := io.ReadFull(exeFile, header.NameBytes); err != nil {
		return "", nil, fmt.Errorf("failed to read name: %w", err)
	}

	header.ScriptBytes = make([]byte, header.ScriptSize)
	if _, err := io.ReadFull(exeFile, header.ScriptBytes); err != nil {
		return "", nil, fmt.Errorf("failed to read script: %w", err)
	}

	actualChecksum := sha256.Sum256(header.ScriptBytes)
	if !bytes.Equal(actualChecksum[:], header.Checksum[:]) {
		return "", nil, fmt.Errorf("checksum mismatch - embedded script may be corrupted")
	}

	scriptName := string(header.NameBytes)
	return scriptName, header.ScriptBytes, nil
}

func HasEmbeddedScript() bool {
	_, _, err := ExtractEmbeddedScript()
	return err == nil
}

func GetEmbeddedBinaryExtension() string {
	if runtime.GOOS == "windows" {
		return ".exe"
	}
	return ""
}
