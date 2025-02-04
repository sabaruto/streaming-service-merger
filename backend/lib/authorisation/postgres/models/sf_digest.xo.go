package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// DigestByTextAndText calls the stored function 'public.digest(text, text) bytea' on db.
func DigestByTextAndText(ctx context.Context, db DB, p0, p1 string) ([]byte, error) {
	// call public.digest
	const sqlstr = `SELECT * FROM public.digest($1, $2)`
	// run
	var r0 []byte
	logf(sqlstr, p0, p1)
	if err := db.QueryRowContext(ctx, sqlstr, p0, p1).Scan(&r0); err != nil {
		return nil, logerror(err)
	}
	return r0, nil
}

// DigestByByteaAndText calls the stored function 'public.digest(bytea, text) bytea' on db.
func DigestByByteaAndText(ctx context.Context, db DB, p0 []byte, p1 string) ([]byte, error) {
	// call public.digest
	const sqlstr = `SELECT * FROM public.digest($1, $2)`
	// run
	var r0 []byte
	logf(sqlstr, p0, p1)
	if err := db.QueryRowContext(ctx, sqlstr, p0, p1).Scan(&r0); err != nil {
		return nil, logerror(err)
	}
	return r0, nil
}
