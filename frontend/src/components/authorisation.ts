import { NextResponse, NextRequest } from "next/server";
import { PartialMiddleware } from "@/common/types";


export const authware: PartialMiddleware = {
  name: 'Authorisation',
  // TODO: Actually check the authorisation status of the request
  matched: (request: NextRequest) => {
    return request.nextUrl.pathname === '/';
  },

  run: (request: NextRequest) => {

    return NextResponse.next();
    // return NextResponse.redirect(new URL('/login', request.url))
  }
}
