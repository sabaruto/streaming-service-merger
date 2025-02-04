package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// HmacByTextTextAndText calls the stored function 'public.hmac(text, text, text) bytea' on db.
func HmacByTextTextAndText(ctx context.Context, db DB, p0, p1, p2 string) ([]byte, error) {
	// call public.hmac
	const sqlstr = `SELECT * FROM public.hmac($1, $2, $3)`
	// run
	var r0 []byte
	logf(sqlstr, p0, p1, p2)
	if err := db.QueryRowContext(ctx, sqlstr, p0, p1, p2).Scan(&r0); err != nil {
		return nil, logerror(err)
	}
	return r0, nil
}

// HmacByByteaByteaAndText calls the stored function 'public.hmac(bytea, bytea, text) bytea' on db.
func HmacByByteaByteaAndText(ctx context.Context, db DB, p0, p1 []byte, p2 string) ([]byte, error) {
	// call public.hmac
	const sqlstr = `SELECT * FROM public.hmac($1, $2, $3)`
	// run
	var r0 []byte
	logf(sqlstr, p0, p1, p2)
	if err := db.QueryRowContext(ctx, sqlstr, p0, p1, p2).Scan(&r0); err != nil {
		return nil, logerror(err)
	}
	return r0, nil
}
