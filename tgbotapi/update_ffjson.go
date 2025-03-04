// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: update.go

package tgbotapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *Update) MarshalJSON() ([]byte, error) {
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
func (j *Update) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"update_id":`)
	fflib.FormatBits2(buf, uint64(j.UpdateID), 10, j.UpdateID < 0)
	if j.Message != nil {
		buf.WriteString(`,"message":`)

		{

			err = j.Message.MarshalJSONBuf(buf)
			if err != nil {
				return err
			}

		}
	} else {
		buf.WriteString(`,"message":null`)
	}
	if j.EditedMessage != nil {
		buf.WriteString(`,"edited_message":`)

		{

			err = j.EditedMessage.MarshalJSONBuf(buf)
			if err != nil {
				return err
			}

		}
	} else {
		buf.WriteString(`,"edited_message":null`)
	}
	if j.ChannelPost != nil {
		buf.WriteString(`,"channel_post":`)

		{

			err = j.ChannelPost.MarshalJSONBuf(buf)
			if err != nil {
				return err
			}

		}
	} else {
		buf.WriteString(`,"channel_post":null`)
	}
	if j.EditedChannelPost != nil {
		buf.WriteString(`,"edited_channel_post":`)

		{

			err = j.EditedChannelPost.MarshalJSONBuf(buf)
			if err != nil {
				return err
			}

		}
	} else {
		buf.WriteString(`,"edited_channel_post":null`)
	}
	if j.InlineQuery != nil {
		buf.WriteString(`,"inline_query":`)

		{

			err = j.InlineQuery.MarshalJSONBuf(buf)
			if err != nil {
				return err
			}

		}
	} else {
		buf.WriteString(`,"inline_query":null`)
	}
	if j.ChosenInlineResult != nil {
		/* Struct fall back. type=tgbotapi.ChosenInlineResult kind=struct */
		buf.WriteString(`,"chosen_inline_result":`)
		err = buf.Encode(j.ChosenInlineResult)
		if err != nil {
			return err
		}
	} else {
		buf.WriteString(`,"chosen_inline_result":null`)
	}
	if j.CallbackQuery != nil {
		buf.WriteString(`,"callback_query":`)

		{

			err = j.CallbackQuery.MarshalJSONBuf(buf)
			if err != nil {
				return err
			}

		}
	} else {
		buf.WriteString(`,"callback_query":null`)
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtUpdatebase = iota
	ffjtUpdatenosuchkey

	ffjtUpdateUpdateID

	ffjtUpdateMessage

	ffjtUpdateEditedMessage

	ffjtUpdateChannelPost

	ffjtUpdateEditedChannelPost

	ffjtUpdateInlineQuery

	ffjtUpdateChosenInlineResult

	ffjtUpdateCallbackQuery
)

var ffjKeyUpdateUpdateID = []byte("update_id")

var ffjKeyUpdateMessage = []byte("message")

var ffjKeyUpdateEditedMessage = []byte("edited_message")

var ffjKeyUpdateChannelPost = []byte("channel_post")

var ffjKeyUpdateEditedChannelPost = []byte("edited_channel_post")

var ffjKeyUpdateInlineQuery = []byte("inline_query")

var ffjKeyUpdateChosenInlineResult = []byte("chosen_inline_result")

var ffjKeyUpdateCallbackQuery = []byte("callback_query")

// UnmarshalJSON umarshall json - template of ffjson
func (j *Update) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *Update) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtUpdatebase
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
				currentKey = ffjtUpdatenosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffjKeyUpdateChannelPost, kn) {
						currentKey = ffjtUpdateChannelPost
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyUpdateChosenInlineResult, kn) {
						currentKey = ffjtUpdateChosenInlineResult
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyUpdateCallbackQuery, kn) {
						currentKey = ffjtUpdateCallbackQuery
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyUpdateEditedMessage, kn) {
						currentKey = ffjtUpdateEditedMessage
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyUpdateEditedChannelPost, kn) {
						currentKey = ffjtUpdateEditedChannelPost
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyUpdateInlineQuery, kn) {
						currentKey = ffjtUpdateInlineQuery
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffjKeyUpdateMessage, kn) {
						currentKey = ffjtUpdateMessage
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffjKeyUpdateUpdateID, kn) {
						currentKey = ffjtUpdateUpdateID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyUpdateCallbackQuery, kn) {
					currentKey = ffjtUpdateCallbackQuery
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyUpdateChosenInlineResult, kn) {
					currentKey = ffjtUpdateChosenInlineResult
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyUpdateInlineQuery, kn) {
					currentKey = ffjtUpdateInlineQuery
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyUpdateEditedChannelPost, kn) {
					currentKey = ffjtUpdateEditedChannelPost
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyUpdateChannelPost, kn) {
					currentKey = ffjtUpdateChannelPost
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyUpdateEditedMessage, kn) {
					currentKey = ffjtUpdateEditedMessage
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyUpdateMessage, kn) {
					currentKey = ffjtUpdateMessage
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyUpdateUpdateID, kn) {
					currentKey = ffjtUpdateUpdateID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtUpdatenosuchkey
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

				case ffjtUpdateUpdateID:
					goto handle_UpdateID

				case ffjtUpdateMessage:
					goto handle_Message

				case ffjtUpdateEditedMessage:
					goto handle_EditedMessage

				case ffjtUpdateChannelPost:
					goto handle_ChannelPost

				case ffjtUpdateEditedChannelPost:
					goto handle_EditedChannelPost

				case ffjtUpdateInlineQuery:
					goto handle_InlineQuery

				case ffjtUpdateChosenInlineResult:
					goto handle_ChosenInlineResult

				case ffjtUpdateCallbackQuery:
					goto handle_CallbackQuery

				case ffjtUpdatenosuchkey:
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

handle_UpdateID:

	/* handler: j.UpdateID type=int kind=int quoted=false*/

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

			j.UpdateID = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Message:

	/* handler: j.Message type=tgbotapi.Message kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.Message = nil

		} else {

			if j.Message == nil {
				j.Message = new(Message)
			}

			err = j.Message.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_EditedMessage:

	/* handler: j.EditedMessage type=tgbotapi.Message kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.EditedMessage = nil

		} else {

			if j.EditedMessage == nil {
				j.EditedMessage = new(Message)
			}

			err = j.EditedMessage.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ChannelPost:

	/* handler: j.ChannelPost type=tgbotapi.Message kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.ChannelPost = nil

		} else {

			if j.ChannelPost == nil {
				j.ChannelPost = new(Message)
			}

			err = j.ChannelPost.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_EditedChannelPost:

	/* handler: j.EditedChannelPost type=tgbotapi.Message kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.EditedChannelPost = nil

		} else {

			if j.EditedChannelPost == nil {
				j.EditedChannelPost = new(Message)
			}

			err = j.EditedChannelPost.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_InlineQuery:

	/* handler: j.InlineQuery type=tgbotapi.InlineQuery kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.InlineQuery = nil

		} else {

			if j.InlineQuery == nil {
				j.InlineQuery = new(InlineQuery)
			}

			err = j.InlineQuery.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ChosenInlineResult:

	/* handler: j.ChosenInlineResult type=tgbotapi.ChosenInlineResult kind=struct quoted=false*/

	{
		/* Falling back. type=tgbotapi.ChosenInlineResult kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.ChosenInlineResult)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CallbackQuery:

	/* handler: j.CallbackQuery type=tgbotapi.CallbackQuery kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.CallbackQuery = nil

		} else {

			if j.CallbackQuery == nil {
				j.CallbackQuery = new(CallbackQuery)
			}

			err = j.CallbackQuery.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
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
