import type { Metadata } from "next";
import { DM_Sans } from "next/font/google";
import '../stylesheets/main.css'

const dmSans = DM_Sans({
    subsets: ['latin'],
})
export const metadata: Metadata = {
    title: "Service Streaming Manager",
    description: "A platform to store and access all your music",
};

export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }>) {
    return (
        <html lang="en">
            <body className={`${dmSans.style.fontFamily}, login`}>
                {children}
            </body>
        </html>
    );
}
