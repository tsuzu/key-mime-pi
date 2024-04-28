package hid

import (
	"fmt"
	"os"
)

type KeyboardHID interface {
	WriteKeyboardReport(controlChars int, hidKeyCode int) error
	Close() error
}

func NewKeyboardHID(path string) (KeyboardHID, error) {
	fp, err := os.OpenFile(path, os.O_WRONLY, 0600)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", path, err)
	}

	return &KeyboardHIDImpl{fp: fp}, nil
}

type KeyboardHIDImpl struct {
	fp *os.File
}

func (k *KeyboardHIDImpl) WriteKeyboardReport(controlChars int, hidKeyCode int) error {
	buf := make([]byte, 8)
	buf[0] = byte(controlChars)
	buf[2] = byte(hidKeyCode)
	err := k.writeWithReopen(buf)
	if err != nil {
		return fmt.Errorf("failed to write keyboard report: %w", err)
	}
	buf = make([]byte, 8)
	err = k.writeWithReopen(buf)
	if err != nil {
		return fmt.Errorf("failed to write empty keyboard report: %w", err)
	}

	return nil
}

func (k *KeyboardHIDImpl) writeWithReopen(b []byte) error {
	_, err := k.fp.Write(b)
	if err == nil {
		return nil
	}

	err = k.fp.Close()
	if err != nil {
		return fmt.Errorf("failed to close file: %w", err)
	}

	fp, err := os.OpenFile(k.fp.Name(), os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", k.fp.Name(), err)
	}
	k.fp = fp

	_, err = k.fp.Write(b)
	if err != nil {
		return fmt.Errorf("failed to write keyboard report: %w", err)
	}

	return nil
}

func (k *KeyboardHIDImpl) Close() error {
	return k.fp.Close()
}
