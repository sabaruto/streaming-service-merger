import {Theme} from "@react-navigation/native";
import {Platform} from "react-native";
import {Style} from "@/constants/Styles";

const WEB_FONT_STACK = '"DM Sans", sans-serif'

export type ThemedFns<T extends Style> = (theme: Theme) => T

export type FontStyle = {
    fontFamily: string;
    fontWeight:
        | 'normal'
        | 'bold'
        | '100'
        | '200'
        | '300'
        | '400'
        | '500'
        | '600'
        | '700'
        | '800'
        | '900';
};

export type FontSize = "regular" | "medium" | "bold" | "heavy"

type Font = {
    regular: FontStyle
    medium: FontStyle
    bold: FontStyle
    heavy: FontStyle
}

const fonts: Font = Platform.select({
    web: {
        regular: {
            fontFamily: WEB_FONT_STACK,
            fontWeight: 'normal',
        },
        medium: {
            fontFamily: WEB_FONT_STACK,
            fontWeight: 'normal',
        },
        bold: {
            fontFamily: WEB_FONT_STACK,
            fontWeight: '700',
        },
        heavy: {
            fontFamily: WEB_FONT_STACK,
            fontWeight: '900',
        },
    },
    default: {
        regular: {
            fontFamily: 'DM-Sans',
            fontWeight: 'normal',
        },
        medium: {
            fontFamily: 'DM-Sans',
            fontWeight: 'normal',
        },
        bold: {
            fontFamily: 'DM-Sans',
            fontWeight: '700',
        },
        heavy: {
            fontFamily: 'DM-Sans',
            fontWeight: '900',
        },
    }
})

const defaultTheme: Theme = {
    dark: false,
    colors: {
        primary: "#DCEDF8",
        background: "#FFFFFF",
        card: "#0091AD",
        text: "#000000",
        border: "#4C3B4D",
        notification: "#F26430"
    },
    fonts,
}

const darkTheme: Theme = {
    dark: true,
    colors: {
        primary: "#103A57",
        background: "#000000",
        card: "#026C7C",
        text: "#FFFFFF",
        border: "#B37BA4",
        notification: "#D99AC5",
    },
    fonts,
}

export const Themes: { [key: string]: Theme } = {
    "light": defaultTheme,
    "dark": darkTheme,
}

export interface ThemeProps<T extends Style> {
    theme: Theme
    themedFns?: ThemedFns<T> | ThemedFns<T>[]
}

export function themedFontSizeStyle(theme: Theme, fontSize: FontSize) {
    switch (fontSize) {
        case("regular"):
            return theme.fonts.regular
        case("medium"):
            return theme.fonts.medium
        case("bold"):
            return theme.fonts.bold
        case("heavy"):
            return theme.fonts.heavy
    }
}

export function themedFuncToStyle<T extends Style>(theme: Theme, styleFns: ThemedFns<T> | ThemedFns<T>[]) {
    let returnStyles = []
    if (Array.isArray(styleFns)) {
        for (let func of styleFns) {
            returnStyles.push(func(theme))
        }
    } else {
        returnStyles = [styleFns(theme)]
    }
    return returnStyles
}
