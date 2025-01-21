import { NextRequest, NextResponse } from "next/server";


export type PartialMiddleware = {
    name: string;
    matched: (request: NextRequest) => boolean;
    run(request: NextRequest): NextResponse
}