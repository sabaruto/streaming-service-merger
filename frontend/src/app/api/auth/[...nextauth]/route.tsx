import NextAuth from "next-auth";
import SpotifyProvider from "next-auth/providers/spotify";


const handler = NextAuth({
  providers: [
    SpotifyProvider({
      clientId: process.env.NEXT_PUBLIC_SPOTIFY_CLIENT_ID!,
      clientSecret: process.env.NEXT_PUBLIC_SPOTIFY_CLIENT_SECRET!
    })
  ]
})

export { handler as GET, handler as POST }
