find . -name "*_ffjson.go" -delete
ffjson payment.go
ffjson succesfull_payment.go
ffjson refunded_payment.go
ffjson callback_query.go
ffjson config_create_invoice_link.go
ffjson config_reply_parameters.go
ffjson config_send_invoice.go
ffjson configs.go
ffjson inline_query.go
ffjson inline_query_result.go
ffjson input_message_content.go
ffjson link_preview_options.go
ffjson message.go
ffjson message_origin.go
ffjson msg_gift_info.go
ffjson msg_write_access_allowed.go
ffjson order_info.go
ffjson photo_size.go
ffjson shipping_address.go
ffjson types.go
ffjson update.go
ffjson users_shared.go
go test ./...