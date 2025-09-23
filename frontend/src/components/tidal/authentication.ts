'use client'

import { credentialsProvider, init } from "@tidal-music/auth"

export async function init_tidal_sdk() {
  const clientId = process.env.NEXT_PUBLIC_TIDAL_CLIENT_ID
  const clientSecret = process.env.NEXT_PUBLIC_TIDAL_CLIENT_SECRET
  const credentialsStorageKey = 'clientCredentials'

  if (clientId == undefined || clientSecret == undefined) {
    throw new Error("Tidal credentials Environment variable not set")
  }

  await init({ clientId, clientSecret, credentialsStorageKey });
}

export async function get_auth_token() {
  const credentials = await credentialsProvider.getCredentials()

  if (credentials.token == undefined) {
    throw new Error("Token credentials not created")
  }

  return credentials.token
}
