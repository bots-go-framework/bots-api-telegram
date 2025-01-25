// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: photo_size.go

package tgbotapi

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *PhotoSize) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *PhotoSize) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "file_id":`)
	fflib.WriteJsonString(buf, string(j.FileID))
	buf.WriteByte(',')
	if len(j.FileUniqueID) != 0 {
		buf.WriteString(`"file_unique_id":`)
		fflib.WriteJsonString(buf, string(j.FileUniqueID))
		buf.WriteByte(',')
	}
	if j.Width != 0 {
		buf.WriteString(`"width":`)
		fflib.FormatBits2(buf, uint64(j.Width), 10, j.Width < 0)
		buf.WriteByte(',')
	}
	if j.Height != 0 {
		buf.WriteString(`"height":`)
		fflib.FormatBits2(buf, uint64(j.Height), 10, j.Height < 0)
		buf.WriteByte(',')
	}
	if j.FileSize != 0 {
		buf.WriteString(`"file_size":`)
		fflib.FormatBits2(buf, uint64(j.FileSize), 10, j.FileSize < 0)
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtPhotoSizebase = iota
	ffjtPhotoSizenosuchkey

	ffjtPhotoSizeFileID

	ffjtPhotoSizeFileUniqueID

	ffjtPhotoSizeWidth

	ffjtPhotoSizeHeight

	ffjtPhotoSizeFileSize
)

var ffjKeyPhotoSizeFileID = []byte("file_id")

var ffjKeyPhotoSizeFileUniqueID = []byte("file_unique_id")

var ffjKeyPhotoSizeWidth = []byte("width")

var ffjKeyPhotoSizeHeight = []byte("height")

var ffjKeyPhotoSizeFileSize = []byte("file_size")

// UnmarshalJSON umarshall json - template of ffjson
func (j *PhotoSize) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *PhotoSize) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtPhotoSizebase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtPhotoSizenosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'f':

					if bytes.Equal(ffjKeyPhotoSizeFileID, kn) {
						currentKey = ffjtPhotoSizeFileID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyPhotoSizeFileUniqueID, kn) {
						currentKey = ffjtPhotoSizeFileUniqueID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyPhotoSizeFileSize, kn) {
						currentKey = ffjtPhotoSizeFileSize
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'h':

					if bytes.Equal(ffjKeyPhotoSizeHeight, kn) {
						currentKey = ffjtPhotoSizeHeight
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffjKeyPhotoSizeWidth, kn) {
						currentKey = ffjtPhotoSizeWidth
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyPhotoSizeFileSize, kn) {
					currentKey = ffjtPhotoSizeFileSize
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyPhotoSizeHeight, kn) {
					currentKey = ffjtPhotoSizeHeight
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyPhotoSizeWidth, kn) {
					currentKey = ffjtPhotoSizeWidth
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyPhotoSizeFileUniqueID, kn) {
					currentKey = ffjtPhotoSizeFileUniqueID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyPhotoSizeFileID, kn) {
					currentKey = ffjtPhotoSizeFileID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtPhotoSizenosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtPhotoSizeFileID:
					goto handle_FileID

				case ffjtPhotoSizeFileUniqueID:
					goto handle_FileUniqueID

				case ffjtPhotoSizeWidth:
					goto handle_Width

				case ffjtPhotoSizeHeight:
					goto handle_Height

				case ffjtPhotoSizeFileSize:
					goto handle_FileSize

				case ffjtPhotoSizenosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_FileID:

	/* handler: j.FileID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.FileID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_FileUniqueID:

	/* handler: j.FileUniqueID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.FileUniqueID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Width:

	/* handler: j.Width type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.Width = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Height:

	/* handler: j.Height type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.Height = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_FileSize:

	/* handler: j.FileSize type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.FileSize = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
