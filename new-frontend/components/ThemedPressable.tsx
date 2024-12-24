import {
    Pressable,
    PressableProps,
    PressableStateCallbackType,
    StyleProp,
    StyleSheet,
    TextStyle,
    ViewStyle
} from "react-native";
import {themedFuncToStyle, ThemeProps} from "@/constants/Themes";
import {Theme} from "@react-navigation/native";
import React from "react";

type PressableStyle = ViewStyle | ((state: PressableStateCallbackType) => StyleProp<ViewStyle>)
type ThemedPressableProps<T extends PressableStyle> = PressableProps & ThemeProps<T>
type ThemedButtonProps<T extends PressableStyle> = ThemedPressableProps<T> & {
    text: string,
    textStyle?: TextStyle | ((state: PressableStateCallbackType) => StyleProp<TextStyle>)
}

function PressableFnsStyle<
    T extends StyleProp<ViewStyle>
        | StyleProp<TextStyle>
>
(style: T | ((state: PressableStateCallbackType) => T)): ((state: PressableStateCallbackType) => T) {
    if (typeof style === "function") {
        return style
    } else {
        return function (state: PressableStateCallbackType) {
            return style
        }
    }
}

function defaultPressable(theme: Theme): ViewStyle {
    return {
        borderStyle: "solid",
        borderRadius: 5,
        transitionDuration: "0.2s",
        backgroundColor: theme.colors.card,
    }
}

export function buttonText(theme: Theme): TextStyle {
    return {
        color: theme.colors.background,
        textTransform: "uppercase",
    }
}

export function ThemedPressable<T extends PressableStyle>({
                                                              theme,
                                                              style,
                                                              themedFns,
                                                              ...otherProps
                                                          }: ThemedPressableProps<T>
) {
    const themedStyle = themedFns ? themedFuncToStyle(theme, themedFns) : []
    const defaultStyle = themedFuncToStyle(theme, defaultPressable)

    const themedStylesheet = PressableFnsStyle(StyleSheet.flatten(themedStyle))
    const instanceStylesheet = PressableFnsStyle(StyleSheet.flatten(style))
    const defaultStylesheet = PressableFnsStyle(StyleSheet.flatten(defaultStyle))


    return <Pressable style={(state: PressableStateCallbackType) => [

        instanceStylesheet(state),
        defaultStylesheet(state),
        themedStylesheet(state),
    ]} {...otherProps}/>
}
