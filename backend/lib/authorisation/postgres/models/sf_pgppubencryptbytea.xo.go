package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// PgpPubEncryptByteaByByteaAndBytea calls the stored function 'public.pgp_pub_encrypt_bytea(bytea, bytea) bytea' on db.
func PgpPubEncryptByteaByByteaAndBytea(ctx context.Context, db DB, p0, p1 []byte) ([]byte, error) {
	// call public.pgp_pub_encrypt_bytea
	const sqlstr = `SELECT * FROM public.pgp_pub_encrypt_bytea($1, $2)`
	// run
	var r0 []byte
	logf(sqlstr, p0, p1)
	if err := db.QueryRowContext(ctx, sqlstr, p0, p1).Scan(&r0); err != nil {
		return nil, logerror(err)
	}
	return r0, nil
}

// PgpPubEncryptByteaByByteaByteaAndText calls the stored function 'public.pgp_pub_encrypt_bytea(bytea, bytea, text) bytea' on db.
func PgpPubEncryptByteaByByteaByteaAndText(ctx context.Context, db DB, p0, p1 []byte, p2 string) ([]byte, error) {
	// call public.pgp_pub_encrypt_bytea
	const sqlstr = `SELECT * FROM public.pgp_pub_encrypt_bytea($1, $2, $3)`
	// run
	var r0 []byte
	logf(sqlstr, p0, p1, p2)
	if err := db.QueryRowContext(ctx, sqlstr, p0, p1, p2).Scan(&r0); err != nil {
		return nil, logerror(err)
	}
	return r0, nil
}
