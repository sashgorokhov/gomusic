package utils

import (
	"errors"
	"github.com/cheggaaa/pb"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
)

func Download_size(url string) (uint64, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, errors.New("HEAD " + url + " " + resp.Status)
	}
	content_length := resp.Header.Get("Content-Length")
	if content_length == "" {
		return 0, nil
	}
	size, err := strconv.Atoi(content_length)
	if err != nil {
		return 0, err
	}
	return uint64(size), nil
}

type CallbackReader struct {
	total, done uint64
	callback    func(total, done uint64)
	reader      io.Reader
}

func (r *CallbackReader) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	r.done += uint64(n)

	if err == nil && r.callback != nil {
		r.callback(r.total, r.done)
	}
	return n, err
}

func Download_file(url string, filename string) error {
	total, err := Download_size(url)
	if err != nil {
		return err
	}
	_, file := path.Split(filename)

	bar := pb.New(int(total))
	bar.Prefix(file)
	bar.SetUnits(pb.U_BYTES)
	bar.Start()

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, bar.NewProxyReader(resp.Body))
	bar.Finish()
	return err
}
