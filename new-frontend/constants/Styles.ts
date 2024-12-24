import {ImageStyle, Platform, PressableStateCallbackType, StyleProp, TextStyle, ViewStyle} from "react-native";
import {Theme} from "@react-navigation/native";

export type Style =
    TextStyle
    | ViewStyle
    | ImageStyle
    | ((state: PressableStateCallbackType) => StyleProp<ViewStyle>)

export function isViewStyle(style: Style): style is ViewStyle {
    return (style as ViewStyle).animationName !== undefined
}

// Variables to work based off
// - Width
// - Height
// - Theme

export const fullScreen: ViewStyle = {
    display: "flex",
    flex: 1,
}
export const column: ViewStyle = {
    display: "flex",
    flexDirection: "column",
}

export const body: ViewStyle = {
    margin: 0,
}

export const root: ViewStyle = {
    justifyContent: "flex-start",
}

/* Default Variables */
export function primaryBackground(theme: Theme): ViewStyle {
    return {
        backgroundColor: theme.colors.primary
    }
}

export const footer: ViewStyle = {
    flex: 1,
    justifyContent: "flex-end"
}

/* Font Sizes */
export const h1: TextStyle = {
    fontSize: 30,
    textAlign: "center",
}

export const center: ViewStyle = {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
}

console.log("Platform:", Platform.OS)

export function input(theme: Theme): TextStyle {
    return {
        borderStyle: "solid",
        height: 25,
        borderRadius: 5,
        marginTop: 5,

        paddingLeft: 5,
        paddingRight: 5,
        borderColor: theme.colors.primary,
        borderWidth: 2,
        paddingBottom: 3,
        paddingTop: 3,
    }
}

export function button(theme: Theme): ViewStyle {
    return {
        borderStyle: "solid",
        borderRadius: 5,
        borderColor: theme.colors.background,
    }
}

export function mainCard(theme: Theme): ViewStyle {
    return {
        width: 300,
        height: 450,

        padding: 50,
        borderRadius: 15,

        alignItems: "stretch",
        backgroundColor: theme.colors.background,
    }
}

export function defaultHeader(theme: Theme): TextStyle {
    return {
        borderBottomColor: theme.colors.primary,
        borderStyle: "solid",
        borderBottomWidth: 2,
    }
}

export const largeBorder: ViewStyle = {
    display: "flex",
    borderRadius: 15,
}

export const form: ViewStyle = {
    flexDirection: "column",
    marginTop: 25,
    gap: 10,
}

export const formButton: ViewStyle = {
    marginTop: 20,
    width: 75,
    alignSelf: "center",
}
