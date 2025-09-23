'use client'

import { createAPIClient } from "@tidal-music/api"
import { credentialsProvider } from "@tidal-music/auth"

var apiClient: any | undefined = undefined;

export function getApiClient() {
  console.log(credentialsProvider.getCredentials.toString())
  if (apiClient == undefined) {
    apiClient = createAPIClient(credentialsProvider)
  }

  return apiClient
}

export async function search(id: string) {
  const client = getApiClient()
  const { data, error } = await client.GET('/searchResults/{id}', {
    params: {
      path: { id },
      query: { countryCode: 'NO', include: 'topHits' },
    },
  });

  console.log(data)
  console.log("We hit something")
}


