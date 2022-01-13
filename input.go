package cli

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"golang.org/x/term"
)

func ReadInputBool(msg string) (bool, error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return false, errors.New("stdin is not a terminal")
	}

	fmt.Fprint(os.Stdout, msg+": ")
	ans, err := readline(os.Stdin)
	if err != nil {
		return false, err
	}

	v, err := strconv.ParseBool(ans)
	if err != nil {
		return false, err
	}

	return v, nil
}

func ReadPasswordBool(msg string) (bool, error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return false, errors.New("stdin is not a terminal")
	}

	fmt.Fprint(os.Stdout, msg+": ")
	ans, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Fprintln(os.Stdout)
	if err != nil {
		return false, err
	}

	v, err := strconv.ParseBool(string(ans))
	if err != nil {
		return false, err
	}

	return v, nil
}

func ReadInputString(msg string) (string, error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return "", errors.New("stdin is not a terminal")
	}

	fmt.Fprint(os.Stdout, msg+": ")
	ans, err := readline(os.Stdin)
	if err != nil {
		return "", err
	}

	return ans, nil
}

func ReadPasswordString(msg string) (string, error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return "", errors.New("stdin is not a terminal")
	}

	fmt.Fprint(os.Stdout, msg+": ")
	ans, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Fprintln(os.Stdout)
	if err != nil {
		return "", err
	}

	return string(ans), nil
}

func ReadInputInt(msg string) (int, error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return 0, errors.New("stdin is not a terminal")
	}

	fmt.Fprint(os.Stdout, msg+": ")
	ans, err := readline(os.Stdin)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(ans, 10, 64)
	if err != nil {
		return 0, err
	}

	return int(v), nil
}

func ReadPasswordInt(msg string) (int, error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return 0, errors.New("stdin is not a terminal")
	}

	fmt.Fprint(os.Stdout, msg+": ")
	ans, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Fprintln(os.Stdout)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(string(ans), 10, 64)
	if err != nil {
		return 0, err
	}

	return int(v), nil
}

func ReadInputInt32(msg string) (int32, error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return 0, errors.New("stdin is not a terminal")
	}

	fmt.Fprint(os.Stdout, msg+": ")
	ans, err := readline(os.Stdin)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(ans, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(v), nil
}

func ReadPasswordInt32(msg string) (int32, error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return 0, errors.New("stdin is not a terminal")
	}

	fmt.Fprint(os.Stdout, msg+": ")
	ans, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Fprintln(os.Stdout)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(string(ans), 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(v), nil
}

func ReadInputInt64(msg string) (int64, error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return 0, errors.New("stdin is not a terminal")
	}

	fmt.Fprint(os.Stdout, msg+": ")
	ans, err := readline(os.Stdin)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(ans, 10, 64)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func ReadPasswordInt64(msg string) (int64, error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return 0, errors.New("stdin is not a terminal")
	}

	fmt.Fprint(os.Stdout, msg+": ")
	ans, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Fprintln(os.Stdout)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(string(ans), 10, 64)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func ReadInputFloat32(msg string) (float32, error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return 0, errors.New("stdin is not a terminal")
	}

	fmt.Fprint(os.Stdout, msg+": ")
	ans, err := readline(os.Stdin)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseFloat(ans, 32)
	if err != nil {
		return 0, err
	}

	return float32(v), nil
}

func ReadPasswordFloat32(msg string) (float32, error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return 0, errors.New("stdin is not a terminal")
	}

	fmt.Fprint(os.Stdout, msg+": ")
	ans, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Fprintln(os.Stdout)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseFloat(string(ans), 32)
	if err != nil {
		return 0, err
	}

	return float32(v), nil
}

func ReadInputFloat64(msg string) (float64, error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return 0, errors.New("stdin is not a terminal")
	}

	fmt.Fprint(os.Stdout, msg+": ")
	ans, err := readline(os.Stdin)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseFloat(ans, 64)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func ReadPasswordFloat64(msg string) (float64, error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return 0, errors.New("stdin is not a terminal")
	}

	fmt.Fprint(os.Stdout, msg+": ")
	ans, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Fprintln(os.Stdout)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseFloat(string(ans), 32)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func readline(r io.Reader) (string, error) {
	var bytes [1]byte
	var buf []byte

	for {
		size, err := r.Read(bytes[:])
		if size > 0 {
			switch bytes[0] {
			case '\n':
				return string(buf), nil
			case '\r':
				// remove \r from passwords on Windows
			default:
				buf = append(buf, bytes[0])
			}
			continue
		}
		if err != nil {
			if err == io.EOF && len(buf) > 0 {
				return string(buf), nil
			}
			return string(buf), err
		}
	}
}
