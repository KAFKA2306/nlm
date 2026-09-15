package main

import (
	"encoding/binary"
	"errors"
	"fmt"
)

func mp4Duration(data []byte) (string, error) {
	mvhd, err := findBox(data, "mvhd")
	if err != nil {
		return "", err
	}
	if len(mvhd) < 20 {
		return "", errors.New("mvhd box is too short")
	}
	version := mvhd[0]
	var timescale uint32
	var duration uint64
	switch version {
	case 0:
		if len(mvhd) < 20 {
			return "", errors.New("mvhd v0 is too short")
		}
		timescale = binary.BigEndian.Uint32(mvhd[12:16])
		duration = uint64(binary.BigEndian.Uint32(mvhd[16:20]))
	case 1:
		if len(mvhd) < 32 {
			return "", errors.New("mvhd v1 is too short")
		}
		timescale = binary.BigEndian.Uint32(mvhd[20:24])
		duration = binary.BigEndian.Uint64(mvhd[24:32])
	default:
		return "", fmt.Errorf("unsupported mvhd version %d", version)
	}
	if timescale == 0 {
		return "", errors.New("mvhd timescale is zero")
	}
	seconds := int64((duration + uint64(timescale)/2) / uint64(timescale))
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	secs := seconds % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, secs), nil
}

func findBox(data []byte, target string) ([]byte, error) {
	for offset := 0; offset+8 <= len(data); {
		size := uint64(binary.BigEndian.Uint32(data[offset : offset+4]))
		typeName := string(data[offset+4 : offset+8])
		header := uint64(8)
		if size == 1 {
			if offset+16 > len(data) {
				break
			}
			size = binary.BigEndian.Uint64(data[offset+8 : offset+16])
			header = 16
		} else if size == 0 {
			size = uint64(len(data) - offset)
		}
		if size < header || uint64(offset)+size > uint64(len(data)) {
			break
		}
		payloadStart := uint64(offset) + header
		payloadEnd := uint64(offset) + size
		if typeName == target {
			return data[payloadStart:payloadEnd], nil
		}
		if typeName == "moov" {
			if payload, err := findBox(data[payloadStart:payloadEnd], target); err == nil {
				return payload, nil
			}
		}
		offset += int(size)
	}
	return nil, fmt.Errorf("%s box not found", target)
}
