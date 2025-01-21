import { NextResponse, NextRequest } from "next/server";
import { PartialMiddleware } from "./common/types";
import { authware } from "./components/authorisation";

const partials = [
    authware
]

export function middleware(request: NextRequest) {
    for (const partial of partials) {
        if (partial.matched(request)) {
            console.log('Middleware %s matched.', partial.name)
            return partial.run(request)
        }
    }
}